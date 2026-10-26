package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"linkspan/admin/internal/cache"
	"linkspan/admin/internal/common/biz/usercontext"
	"linkspan/admin/internal/common/convention/errorcode"
	"linkspan/admin/internal/common/convention/exception"
	"linkspan/admin/internal/common/serialize"
	"linkspan/admin/internal/dao/entity"
	"linkspan/admin/internal/dto/req"
	"linkspan/admin/internal/dto/resp"
	"linkspan/admin/internal/repository"
	"log"
	"regexp"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repository.UserRepository
	userBloom      *cache.UserBloom
	userLock       *cache.UserRegisterLock
	userSession    *cache.UserSession
}

func NewUserService(userRepository *repository.UserRepository,
	userBloom *cache.UserBloom, userLock *cache.UserRegisterLock,
	usersession *cache.UserSession) *UserService {
	return &UserService{
		userRepository: userRepository,
		userBloom:      userBloom,
		userLock:       userLock,
		userSession:    usersession,
	}
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*resp.UserResp, error) {
	user, err := s.userRepository.SelectByUsername(ctx, username)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, exception.NewServiceError(errorcode.UserNull)
	}

	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	return &resp.UserResp{
		ID:       user.ID,
		Username: user.Username,
		RealName: user.RealName,
		Phone:    serialize.Phone(user.Phone),
		Mail:     user.Mail,
	}, nil
}

// true为可用， false为可能被占用
func (s *UserService) HasUsername(ctx context.Context, username string) (bool, error) {

	exists, err := s.userBloom.Contains(ctx, username)
	if err != nil {
		return false, err
	}

	return !exists, nil
}

var registerUsernamePattern = regexp.MustCompile(
	`^[a-z0-9_]{3,32}$`,
)

func (s *UserService) Register(ctx context.Context, request req.UserRegisterReq) error {
	// 固定用户名格式，并检查 bcrypt 的密码字节限制
	if !registerUsernamePattern.MatchString(request.Username) ||
		len(request.Password) > 72 {
		return exception.NewServiceError(errorcode.ClientError)
	}

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	available, err := s.HasUsername(ctx, request.Username)
	if err != nil {
		return err
	}
	if !available {
		return exception.NewServiceError(errorcode.UserNameExist)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("生成密码哈希失败: %w", err)
	}

	token, acquired, err := s.userLock.TryLock(ctx, request.Username)
	if err != nil {
		return fmt.Errorf("获取注册锁失败: %w", err)
	}

	if !acquired {
		return exception.NewServiceError(errorcode.UserNameExist)
	}

	defer func() {
		unlockCtx, unlockCancel := context.WithTimeout(
			context.Background(),
			3*time.Second,
		)
		defer unlockCancel()

		if err := s.userLock.Unlock(
			unlockCtx,
			request.Username,
			token,
		); err != nil {
			log.Printf("释放用户注册锁失败: %v", err)
		}
	}()

	// 等待和计算期间，其他请求可能已经完成注册
	available, err = s.HasUsername(ctx, request.Username)
	if err != nil {
		return err
	}

	if !available {
		return exception.NewServiceError(errorcode.UserNameExist)
	}

	user := entity.UserDO{
		Username: request.Username,
		Password: string(passwordHash),
		RealName: request.RealName,
		Phone:    request.Phone,
		Mail:     request.Mail,
	}

	err = s.userRepository.WithTransaction(
		ctx,
		func(tx *sql.Tx) error {
			inserted, err := s.userRepository.Insert(
				ctx,
				tx,
				&user,
			)
			if err != nil {
				return err
			}

			if inserted != 1 {
				return exception.NewServiceError(
					errorcode.UserSaveError,
				)
			}

			// 先更新过滤器，成功后才尝试提交MySQL事务
			if err := s.userBloom.Add(
				ctx,
				request.Username,
			); err != nil {
				return err
			}

			return nil
		},
	)

	if err != nil {
		var mysqlErr *mysql.MySQLError

		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return exception.NewServiceError(errorcode.UserExist)
		}

		return fmt.Errorf("注册用户失败: %w", err)
	}

	return nil
}

func (s *UserService) Update(ctx context.Context, request req.UserUpdateReq) error {
	currentUsername, ok := usercontext.Username(ctx)
	if !ok {
		err := exception.NewServiceError(errorcode.ClientError)
		err.Message = "用户未登录或登录已过期"
		return err
	}

	if currentUsername != request.Username {
		err := exception.NewServiceError(errorcode.ClientError)
		err.Message = "当前登录用户修改请求异常"
		return err
	}

	if request.RealName == nil &&
		request.Phone == nil &&
		request.Mail == nil {
		return exception.NewServiceError(errorcode.ClientError)
	}

	update := entity.UserProfileUpdate{
		RealName: request.RealName,
		Phone:    request.Phone,
		Mail:     request.Mail,
	}

	err := s.userRepository.WithTransaction(
		ctx,
		func(tx *sql.Tx) error {
			return s.userRepository.UpdateProfile(
				ctx,
				tx,
				currentUsername,
				update,
			)
		},
	)

	if errors.Is(err, sql.ErrNoRows) {
		return exception.NewServiceError(errorcode.UserNull)
	}

	if err != nil {
		return fmt.Errorf("修改用户资料失败: %w", err)
	}

	return nil
}

func loginFailedError() *exception.ServiceError {
	err := exception.NewServiceError(errorcode.ClientError)
	err.Message = "用户名或密码错误"
	return err
}

func (s *UserService) Login(ctx context.Context, request req.UserLoginReq) (*resp.UserLoginResp, error) {
	if !registerUsernamePattern.MatchString(request.Username) || len(request.Password) > 72 {
		return nil, exception.NewServiceError(errorcode.ClientError)
	}

	user, err := s.userRepository.SelectForLogin(ctx, request.Username)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, loginFailedError()
	}

	if err != nil {
		return nil, fmt.Errorf("查询登录用户失败: %w", err)
	}

	if user.Password == "" {
		return nil, loginFailedError()
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, loginFailedError()
	}

	if err != nil {
		return nil, fmt.Errorf("验证用户密码失败: %w", err)
	}

	token, err := s.userSession.CreateOrRefresh(
		ctx,
		cache.SessionUser{
			ID:       user.ID,
			Username: user.Username,
			RealName: user.RealName,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("保存登录会话失败: %w", err)
	}

	return &resp.UserLoginResp{
		Token: token,
	}, nil
}


func (s *UserService) CheckLogin(ctx context.Context, username string, token string) (bool, error) {
	valid, err := s.userSession.Exists(ctx, username, token)

	if err != nil {
		return false, fmt.Errorf("检查登录状态失败: %w", err)
	}	

	return valid, nil
}

func (s *UserService) Logout(ctx context.Context, username string, token string) error {
	deleted, err := s.userSession.DeleteIfValid(ctx, username, token)
	if err != nil {
		return fmt.Errorf("退出登录失败: %w", err)
	}

	if !deleted {
		err := exception.NewServiceError(errorcode.ClientError)
		err.Message = "用户Token不存在或用户未登录"
		return err
	}

	return nil
}