package mapper

import "database/sql"

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

