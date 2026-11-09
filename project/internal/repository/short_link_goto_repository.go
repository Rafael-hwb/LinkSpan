package repository

import (
	"context"
	"database/sql"
	"fmt"
	"linkspan/project/internal/dao/entity"
)

type ShortLinkGotoRepository struct {
}

func NewShortLinkGotoRepository() *ShortLinkGotoRepository {
	return &ShortLinkGotoRepository{}
}

func (r *ShortLinkGotoRepository) Insert(ctx context.Context, tx *sql.Tx, link entity.ShortLinkGotoDO) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (gid, full_short_url)
		VALUES (?, ?)
	`, ShortLinkGotoTableName(link.FullShortURL))

	if _, err := tx.ExecContext(ctx, query, link.GID, link.FullShortURL); err != nil {
		return fmt.Errorf("保存短链接跳转路由失败: %w", err)
	}

	return nil
}