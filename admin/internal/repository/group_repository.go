package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type GroupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{
		db: db,
	}
}

func (r *GroupRepository) WithTransaction(
	ctx context.Context,
	action func(*sql.Tx) error,
) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return fmt.Errorf("开启分组事务失败: %w", err)
	}

	defer tx.Rollback()

	if err := action(tx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("提交分组事务失败: %w", err)
	}

	return nil
}

func (r *GroupRepository) LockUser(ctx context.Context, tx *sql.Tx, username string) error {
	query := fmt.Sprintf(`
		SELECT id
		FROM %s
		WHERE username = ?
		  AND del_flag = 0
		FOR UPDATE
	`, UserTableName(username))

	var id int64

	if err := tx.QueryRowContext(ctx, query, username).Scan(&id); err != nil {
		return fmt.Errorf("锁定分组所属用户失败: %w", err)
	}

	return nil
}

func (r *GroupRepository) CountByUsername(ctx context.Context, tx *sql.Tx, username string) (int, error) {
	query := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM %s
		WHERE username = ?
		  AND del_flag = 0
	`, GroupTableName(username))
	
	var count int

	if err := tx.QueryRowContext(ctx, query, username).Scan(&count); err != nil {
		return 0, fmt.Errorf("统计用户分组失败: %w", err)
	}

	return count, nil
}

func (r *GroupRepository) ReserveGID(ctx context.Context, tx *sql.Tx, gid string) (bool, error) {
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

func (r *GroupRepository) Insert(ctx context.Context, tx *sql.Tx, username string, gid string, name string) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			gid,
			name,
			username,
			sort_order,
			create_time,
			update_time,
			del_flag
		)
		VALUES (?, ?, ?, 0, NOW(), NOW(), 0)
	`, GroupTableName(username))
	
	_, err := tx.ExecContext(ctx, query, gid, name, username)

	if err != nil {
		return fmt.Errorf("新增分组失败: %w", err)
	}

	return nil
}