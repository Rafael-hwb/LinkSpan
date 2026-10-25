package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type legacyUserFields struct {
	id    int64
	phone sql.NullString
	mail  sql.NullString
}

// MigrateSensitiveFields 用于本地少量测试数据的停机迁移
// 每张表使用独立事务，允许失败修正后重新执行
func (r *UserRepository) MigrateSensitiveFields(
	ctx context.Context,
) error {
	for i := 0; i < UserShardCount; i++ {
		table := fmt.Sprintf("t_user_%d", i)

		err := r.WithTransaction(ctx, func(tx *sql.Tx) error {
			query := fmt.Sprintf(
				"SELECT id, phone, mail FROM %s",
				table,
			)

			rows, err := tx.QueryContext(ctx, query)
			if err != nil {
				return err
			}
			defer rows.Close()

			var users []legacyUserFields

			for rows.Next() {
				var user legacyUserFields

				if err := rows.Scan(
					&user.id,
					&user.phone,
					&user.mail,
				); err != nil {
					return err
				}

				users = append(users, user)
			}

			if err := rows.Err(); err != nil {
				return err
			}

			// 同一事务连接上执行 UPDATE 前，先关闭查询结果
			if err := rows.Close(); err != nil {
				return err
			}

			updateSQL := fmt.Sprintf(
				"UPDATE %s SET phone = ?, mail = ? WHERE id = ?",
				table,
			)

			for _, user := range users {
				phone, err := r.encryptLegacyField(
					"phone",
					user.phone,
				)
				if err != nil {
					return err
				}

				mail, err := r.encryptLegacyField(
					"mail",
					user.mail,
				)
				if err != nil {
					return err
				}

				if _, err := tx.ExecContext(
					ctx,
					updateSQL,
					phone,
					mail,
					user.id,
				); err != nil {
					return err
				}
			}

			return nil
		})

		if err != nil {
			return fmt.Errorf("迁移 %s 失败: %w", table, err)
		}
	}

	return nil
}

func (r *UserRepository) encryptLegacyField(
	field string,
	value sql.NullString,
) (sql.NullString, error) {
	if !value.Valid || value.String == "" {
		return value, nil
	}

	// 已迁移的数据先验证，再保持原值，避免重复加密
	if strings.HasPrefix(value.String, "enc:") {
		if _, err := r.fieldCipher.Decrypt(
			field,
			value.String,
		); err != nil {
			return sql.NullString{}, err
		}

		return value, nil
	}

	encrypted, err := r.fieldCipher.Encrypt(
		field,
		value.String,
	)
	if err != nil {
		return sql.NullString{}, err
	}

	return sql.NullString{
		String: encrypted,
		Valid:  true,
	}, nil
}
