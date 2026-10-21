package mapper

import (
	"context"
	"database/sql"
	"fmt"
	"linkspan/admin/internal/dao/entity"
)

//用户表数据库操作
type UserMapper struct {
	db *sql.DB
}

//创建Mapper并传入数据库连接池
func NewUserMapper(db *sql.DB) *UserMapper {
	return &UserMapper{
		db: db,
	}
}

//根据用户名查找用户
func (m *UserMapper) SelectByUsername(ctx context.Context, username string) (*entity.UserDO, error) {
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
