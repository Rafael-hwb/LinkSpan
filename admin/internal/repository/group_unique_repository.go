package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type GroupUniqueRepository struct {
	db *sql.DB
}

func NewGroupUniqueRepository(db *sql.DB) *GroupUniqueRepository {
	return &GroupUniqueRepository{
		db: db,
	}
}


func (r *GroupUniqueRepository) Reserve(ctx context.Context, tx *sql.Tx, gid string) (bool, error) {
	_, err := tx.ExecContext(ctx, `
		INSERT INTO t_group_unique (gid)
		VALUES (?)
	`, gid)
	if err == nil {
		return true, nil
	}

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return false, nil
	}

	return false, fmt.Errorf("登记分组标识失败: %w", err)	
}

func (r *GroupUniqueRepository) ForEachGID(ctx context.Context, visit func(string) error) error {
	rows, err := r.db.QueryContext(ctx, `
		SELECT gid
		FROM t_group_unique
		WHERE gid IS NOT NULL
	`)
	if err != nil {
		return fmt.Errorf("查询已有分组标识失败: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var gid string
		if err := rows.Scan(&gid); err != nil {
			return fmt.Errorf("读取分组标识失败: %w", err)
		}

		if err := visit(gid); err != nil {
			return err
		}
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("遍历分组标识失败: %w", err)
	}

	return nil
}