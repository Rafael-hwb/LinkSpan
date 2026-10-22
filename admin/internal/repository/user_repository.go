package repository

import (
	"context"
	"database/sql"
	"fmt"
	"linkspan/admin/internal/dao/entity"
)

// 用户表数据库操作
type UserRepository struct {
	db *sql.DB
}

// 创建Repository并传入数据库连接池
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// 根据用户名查找用户
func (m *UserRepository) SelectByUsername(ctx context.Context, username string) (*entity.UserDO, error) {
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

	err := m.db.QueryRowContext(
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

	return &user, nil
}
