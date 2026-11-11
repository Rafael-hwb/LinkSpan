package repository

import (
	"context"
	"database/sql"
	"fmt"
	"linkspan/project/internal/dao/entity"
)

type ShortLinkGotoRepository struct {
	db *sql.DB
}

func NewShortLinkGotoRepository(db *sql.DB) *ShortLinkGotoRepository {
	return &ShortLinkGotoRepository{
		db: db,
	}
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

func (r *ShortLinkGotoRepository) FindByFullShortURL(ctx context.Context, fullShortURL string) (entity.ShortLinkGotoDO, error) {
	query := fmt.Sprintf(`
		SELECT gid, full_short_url
		FROM %s
		WHERE full_short_url = ?
	`, ShortLinkGotoTableName(fullShortURL))

	var route entity.ShortLinkGotoDO
	err := r.db.QueryRowContext(ctx, query, fullShortURL).Scan(&route.GID, &route.FullShortURL)
	if err != nil {
		return entity.ShortLinkGotoDO{}, fmt.Errorf("查询短链接跳转路由失败: %w", err)
	}

	return route, nil
}