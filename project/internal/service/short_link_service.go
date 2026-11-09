package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"linkspan/project/internal/dao/entity"
	"linkspan/project/internal/dto/req"
	"linkspan/project/internal/dto/resp"
	"linkspan/project/internal/repository"

	"github.com/go-sql-driver/mysql"
)

var ErrInvalidCreateRequest = errors.New("创建短链接参数错误")

type ShortLinkService struct {
	links  *repository.ShortLinkRepository
	routes *repository.ShortLinkGotoRepository
}

func NewShortLinkService(links *repository.ShortLinkRepository, routes *repository.ShortLinkGotoRepository) *ShortLinkService {
	return &ShortLinkService{links: links, routes: routes}
}

func (s *ShortLinkService) Create(ctx context.Context, request req.ShortLinkCreateReq) (resp.ShortLinkCreateResp, error) {
	var response resp.ShortLinkCreateResp
	link, err := buildShortLink(request)
	if err != nil {
		return response, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for attempt := 0; attempt < 10; attempt++ {
		var random [4]byte
		if _, err := rand.Read(random[:]); err != nil {
			return response, fmt.Errorf("生成短链接后缀失败: %w", err)
		}
		link.ShortURI = hex.EncodeToString(random[:])
		link.FullShortURL = link.Domain + "/" + link.ShortURI

		route := entity.ShortLinkGotoDO{GID: link.GID, FullShortURL: link.FullShortURL}
		err := s.links.WithTransaction(ctx, func(tx *sql.Tx) error {
			if err := s.links.Insert(ctx, tx, link); err != nil {
				return err
			}
			return s.routes.Insert(ctx, tx, route)
		})
		if err == nil {
			return resp.ShortLinkCreateResp{
				GID: link.GID, OriginURL: link.OriginURL, FullShortURL: "http://" + link.FullShortURL,
			}, nil
		}
		var mysqlErr *mysql.MySQLError
		if !errors.As(err, &mysqlErr) || mysqlErr.Number != 1062 {
			return response, err
		}
	}
	return response, errors.New("短链接标识重试次数已耗尽")
}

func buildShortLink(request req.ShortLinkCreateReq) (entity.ShortLinkDO, error) {
	var link entity.ShortLinkDO
	domain := strings.ToLower(strings.TrimSpace(request.Domain))
	if domain != "localhost:8002" {
		return link, fmt.Errorf("%w: domain 须为 localhost:8002", ErrInvalidCreateRequest)
	}
	origin := strings.TrimSpace(request.OriginURL)
	target, err := url.Parse(origin)
	if err != nil || target.Hostname() == "" || (target.Scheme != "http" && target.Scheme != "https") {
		return link, fmt.Errorf("%w: originUrl 须为完整的 HTTP 或 HTTPS 地址", ErrInvalidCreateRequest)
	}
	gid := strings.TrimSpace(request.GID)
	if gid == "" {
		return link, fmt.Errorf("%w: gid 不能为空", ErrInvalidCreateRequest)
	}
	if request.CreatedType == nil || (*request.CreatedType != 0 && *request.CreatedType != 1) || request.ValidDateType == nil {
		return link, fmt.Errorf("%w: 请检查创建类型和有效期类型", ErrInvalidCreateRequest)
	}

	var validDate *time.Time
	switch *request.ValidDateType {
	case 0:
		if request.ValidDate != nil {
			return link, fmt.Errorf("%w: 永久有效时 validDate 须为 null", ErrInvalidCreateRequest)
		}
	case 1:
		if request.ValidDate == nil {
			return link, fmt.Errorf("%w: 自定义有效期需要 validDate", ErrInvalidCreateRequest)
		}
		const layout = "2006-01-02 15:04:05"
		parsed, err := time.ParseInLocation(layout, *request.ValidDate, time.FixedZone("UTC+8", 8*60*60))
		if err != nil || parsed.Format(layout) != *request.ValidDate || !parsed.After(time.Now()) {
			return link, fmt.Errorf("%w: validDate 须为未来的北京时间，格式为 yyyy-MM-dd HH:mm:ss", ErrInvalidCreateRequest)
		}
		validDate = &parsed
	default:
		return link, fmt.Errorf("%w: validDateType 须为 0 或 1", ErrInvalidCreateRequest)
	}
	return entity.ShortLinkDO{
		Domain: domain, OriginURL: origin, GID: gid, CreatedType: *request.CreatedType,
		ValidDateType: *request.ValidDateType, ValidDate: validDate, Describe: request.Describe,
	}, nil
}
