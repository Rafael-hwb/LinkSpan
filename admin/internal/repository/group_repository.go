package repository

import (
	"context"
	"database/sql"
	"fmt"
	"linkspan/admin/internal/dao/entity"
)

type GroupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{
		db: db,
	}
}

func (r *GroupRepository) CountByUsername(ctx context.Context, username string) (int, error) {
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE username = ? AND del_flag = 0`, GroupTableName(username))
	var count int
	if err := r.db.QueryRowContext(ctx, query, username).Scan(&count); err != nil {
		return 0, fmt.Errorf("统计用户分组失败: %w", err)
	}
	return count, nil
}

func (r *GroupRepository) WithTransaction(ctx context.Context, action func(*sql.Tx) error) error {
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

func (r *GroupRepository) Insert(ctx context.Context, tx *sql.Tx, group entity.GroupDO) error {
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
		VALUES (?, ?, ?, ?, NOW(), NOW(), 0)
	`, GroupTableName(group.Username))

	_, err := tx.ExecContext(ctx, query, group.GID, group.Name, group.Username, group.SortOrder)
	if err != nil {
		return fmt.Errorf("新增分组失败: %w", err)
	}
	return nil
}

func (r *GroupRepository) ListByUsername(ctx context.Context, username string) ([]entity.GroupDO, error) {
	query := fmt.Sprintf(`
		SELECT gid, name, username, COALESCE(sort_order, 0)
		FROM %s
		WHERE username = ?
		  AND del_flag = 0
		ORDER BY sort_order DESC, update_time DESC, id DESC
	`, GroupTableName(username))

	rows, err := r.db.QueryContext(ctx, query, username)
	if err != nil {
		return nil, fmt.Errorf("查询用户分组失败: %w", err)
	}
	defer rows.Close()

	groups := make([]entity.GroupDO, 0)

	for rows.Next() {
		var group entity.GroupDO
		if err := rows.Scan(&group.GID, &group.Name, &group.Username, &group.SortOrder); err != nil {
			return nil, fmt.Errorf("读取分组记录失败: %w", err)
		}

		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历分组记录失败: %w", err)
	}

	return groups, nil
}


func (r *GroupRepository) UpdateName(ctx context.Context, username string, gid string, name string) error {
	groupTableName := GroupTableName(gid)

	query := fmt.Sprintf(`
		UPDATE %s
		SET name = ?, update_time = NOW()
		WHERE username = ?
		  AND gid = ?
		  AND del_flag = 0
	`, groupTableName)

	if _, err := r.db.ExecContext(ctx, query, name, username, gid); err != nil {
		return fmt.Errorf("修改分组名称失败: %w", err) 
	}
	
	return nil
}