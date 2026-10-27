package repository

import (
	"context"
	"database/sql"
	"fmt"
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
