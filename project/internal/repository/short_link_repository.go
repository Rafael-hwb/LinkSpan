package repository

import (
	"context"
	"database/sql"
	"fmt"
	"linkspan/project/internal/dao/entity"
)

type ShortLinkRepository struct {
	db *sql.DB
}

func NewShortLinkRepository(db *sql.DB) ShortLinkRepository {
	return ShortLinkRepository{
		db: db,
	}
}

func (r *ShortLinkRepository) WithTransaction(ctx context.Context, action func(*sql.Tx) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("开启短链接事务失败: %w", err)
	}
	defer tx.Rollback()

	if err := action(tx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("提交短链接事务失败: %w", err)
	}

	return nil
}

func (r *ShortLinkRepository) Insert(ctx context.Context, tx *sql.Tx, link entity.ShortLinkDO) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			domain,
			short_uri,
			full_short_url,
			origin_url,
			click_num,
			gid,
			favicon,
			enable_status,
			created_type,
			valid_date_type,
			valid_date,
			%s,
			total_pv,
			total_uv,
			total_uip,
			create_time,
			update_time,
			del_time,
			del_flag
		)
		VALUES (?, ?, ?, ?, 0, ?, ?, 0, ?, ?, ?, ?, 0, 0, 0, NOW(), NOW(), 0, 0)
	`, ShortLinkTableName(link.GID), "`describe`")

	_, err := tx.ExecContext(
		ctx,
		query,
		link.Domain,
		link.ShortURI,
		link.FullShortURL,
		link.OriginURL,
		link.GID,
		link.Favicon,
		link.CreatedType,
		link.ValidDateType,
		link.ValidDate,
		link.Describe,
	)
	if err != nil {
		return fmt.Errorf("保存短链接失败: %w", err)
	}

	return nil
}


func (r *ShortLinkRepository) PageByGID(ctx context.Context, gid string, limit, offset int64) ([]entity.ShortLinkDO, int64, error) {
	filter := fmt.Sprintf("FROM %s WHERE gid = ? AND enable_status = 0 AND del_flag = 0", ShortLinkTableName(gid))

	var total int64
	if err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) "+filter, gid).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("统计分组短链接失败: %w", err)
	}

	query := "SELECT gid, full_short_url, origin_url, COALESCE(`describe`, '') " + filter +
		" ORDER BY create_time DESC, id DESC LIMIT ? OFFSET ?"

	rows, err := r.db.QueryContext(ctx, query, gid, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("分页查询短链接失败: %w", err)
	}
	defer rows.Close()

	links := make([]entity.ShortLinkDO, 0)
	for rows.Next() {
		var link entity.ShortLinkDO
		if err := rows.Scan(&link.GID, &link.FullShortURL, &link.OriginURL, &link.Describe); err != nil {
			return nil, 0, fmt.Errorf("读取短链接失败: %w", err)
		}
		links = append(links, link)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("遍历短链接失败: %w", err)
	}

	return links, total, nil
}