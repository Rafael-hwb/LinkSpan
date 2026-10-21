package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

// 创建连接池并检查数据库连接
func OpenDatabase() (*sql.DB, error) {
	password, exists := os.LookupEnv("LINKSPAN_DB_PASSWORD")
	if !exists {
		return nil, fmt.Errorf("请设置 LINKSPAN_DB_PASSWORD")
	}

	cfg := mysql.NewConfig()

	cfg.User = os.Getenv("LINKSPAN_DB_USER")
	if cfg.User == "" {
		cfg.User = "root"
	}

	cfg.Passwd = password
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "linkspan"
	cfg.ParseTime = true
	cfg.Timeout = 5 * time.Second
	cfg.ReadTimeout = 5 * time.Second
	cfg.WriteTimeout = 5 * time.Second

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("连接数据库失败: %w", err)
	}

	return db, nil
}