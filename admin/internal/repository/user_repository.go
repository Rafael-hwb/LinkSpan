package repository

import (
	"context"
	"database/sql"
	"fmt"
	"linkspan/admin/internal/common/security"
	"linkspan/admin/internal/dao/entity"
	"strings"
)

// 用户表数据库操作
type UserRepository struct {
	db          *sql.DB
	fieldCipher *security.FieldCipher
}

// 创建Repository并传入数据库连接池
func NewUserRepository(db *sql.DB,
	fieldCipher *security.FieldCipher) *UserRepository {
	return &UserRepository{
		db:          db,
		fieldCipher: fieldCipher,
	}
}

// 根据用户名查找用户
func (r *UserRepository) SelectByUsername(ctx context.Context, username string) (*entity.UserDO, error) {
	table := UserTableName(username)

	query := fmt.Sprintf(`
		SELECT
			id,
			username,
			COALESCE(real_name, ''),
			COALESCE(phone, ''),
			COALESCE(mail, '')
		FROM %s
		WHERE username = ?
		  AND del_flag = 0
		LIMIT 1
	`, table)

	var user entity.UserDO

	err := r.db.QueryRowContext(
		ctx,
		query,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.RealName,
		&user.Phone,
		&user.Mail,
	)

	if err != nil {
		return nil, err
	}

	user.Phone, err = r.fieldCipher.Decrypt(
		"phone",
		user.Phone,
	)
	if err != nil {
		return nil, fmt.Errorf("读取用户手机号失败: %w", err)
	}

	user.Mail, err = r.fieldCipher.Decrypt(
		"mail",
		user.Mail,
	)
	if err != nil {
		return nil, fmt.Errorf("读取用户邮箱失败: %w", err)
	}

	return &user, nil
}

// 判断用户名是否存在
func (r *UserRepository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	table := UserTableName(username)

	query := fmt.Sprintf(`
		SELECT EXISTS(
			SELECT 1
			FROM %s
			WHERE username = ?
		)
		`, table)

	var exists bool

	err := r.db.QueryRowContext(
		ctx,
		query,
		username,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *UserRepository) ForEachUsername(ctx context.Context, visit func(string) error) error {
	for i := 0; i < UserShardCount; i++ {
		query := fmt.Sprintf(`
			SELECT username
			FROM t_user_%d
			WHERE username IS NOT NULL
		`, i)

		if err := r.VisitUsernames(ctx, query, visit); err != nil {
			return err
		}
	}
	return nil
}

func (r *UserRepository) VisitUsernames(ctx context.Context, query string, visit func(string) error) error {
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var username string

		if err := rows.Scan(&username); err != nil {
			return err
		}

		if err := visit(username); err != nil {
			return err
		}
	}
	return rows.Err()
}

// 开启一个mysql事务
func (r *UserRepository) WithTransaction(
	ctx context.Context,
	action func(*sql.Tx) error,
) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := action(tx); err != nil {
		return err
	}
	return tx.Commit()
}

func (r *UserRepository) Insert(ctx context.Context,
	tx *sql.Tx, user *entity.UserDO) (int64, error) {
	encryptedPhone, err := r.fieldCipher.Encrypt("phone", user.Phone)
	if err != nil {
		return 0, err
	}

	encryptedMail, err := r.fieldCipher.Encrypt("mail", user.Mail)
	if err != nil {
		return 0, err
	}

	table := UserTableName(user.Username)

	query := fmt.Sprintf(`
		INSERT INTO %s
		(
			username,
			password,
			real_name,
			phone,
			mail,
			deletion_time,
			create_time,
			update_time,
			del_flag
		)
		VALUES (?, ?, ?, ?, ?, 0, NOW(), NOW(), 0)
	`, table)

	executed, err := tx.ExecContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.RealName,
		encryptedPhone,
		encryptedMail,
	)

	if err != nil {
		return 0, err
	}

	return executed.RowsAffected()
}

func (r *UserRepository) UpdateProfile(ctx context.Context, tx *sql.Tx,
	username string, update entity.UserProfileUpdate) error {
	table := UserTableName(username)

	lockSQL := fmt.Sprintf(`
		SELECT id
		FROM %s
		WHERE username = ? AND del_flag = 0
		FOR UPDATE
	`, table)

	var id int64

	if err := tx.QueryRowContext(ctx, lockSQL, username).Scan(&id); err != nil {
		return err
	}

	var sets []string
	var args []any

	if update.RealName != nil {
		sets = append(sets, "real_name = ?")
		args = append(args, *update.RealName)
	}

	if update.Phone != nil {
		encrypted, err := r.fieldCipher.Encrypt(
			"phone",
			*update.Phone,
		)
		if err != nil {
			return err
		}

		sets = append(sets, "phone = ?")
		args = append(args, encrypted)
	}

	if update.Mail != nil {
		encrypted, err := r.fieldCipher.Encrypt(
			"mail",
			*update.Mail,
		)
		if err != nil {
			return err
		}

		sets = append(sets, "mail = ?")
		args = append(args, encrypted)
	}
	if len(sets) == 0 {
		return fmt.Errorf("没有需要更新的资料字段")
	}

	sets = append(sets, "update_time = NOW()")
	args = append(args, id)

	updateSQL := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = ?",
		table,
		strings.Join(sets, ", "),
	)

	_, err := tx.ExecContext(ctx, updateSQL, args...)
	return err
}

// 读取身份验证所需字段，不读取手机号和邮箱
func (r *UserRepository) SelectForLogin(ctx context.Context, username string) (*entity.UserDO, error) {
	table := UserTableName(username)

	query := fmt.Sprintf(`
		SELECT
			id,
			username,
			COALESCE(password, ''),
			COALESCE(real_name, '')
		FROM %s
		WHERE username = ?
		  AND del_flag = 0
		LIMIT 1
	`, table)

	var user entity.UserDO

	err := r.db.QueryRowContext(
		ctx,
		query,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.RealName,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
