// Package sqlite contains SQLite-specific schema helpers.
package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	sqlcore "github.com/G226Medi03/SQLCore"
)

// Executor is the database surface accepted by store helpers.
type Executor = sqlcore.Executor

// TableExists reports whether a table exists in the current database.
func TableExists(ctx context.Context, db Executor, table string) (bool, error) {
	var name string
	err := db.QueryRowContext(ctx,
		`SELECT name FROM sqlite_master WHERE type = 'table' AND name = ?`, table).Scan(&name)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// EnsureRenamedColumn renames a table column when an older schema still uses
// the old name.
func EnsureRenamedColumn(ctx context.Context, db Executor, table, oldColumn, newColumn string) error {
	rows, err := db.QueryContext(ctx, `PRAGMA table_info(`+table+`)`)
	if err != nil {
		return fmt.Errorf("%s schema 조회 실패: %w", table, err)
	}
	hasOld := false
	hasNew := false
	for rows.Next() {
		var cid, notnull, pk int
		var name, columnType string
		var defaultValue sql.NullString
		if err := rows.Scan(&cid, &name, &columnType, &notnull, &defaultValue, &pk); err != nil {
			_ = rows.Close()
			return err
		}
		switch name {
		case oldColumn:
			hasOld = true
		case newColumn:
			hasNew = true
		}
	}
	if err := rows.Err(); err != nil {
		_ = rows.Close()
		return err
	}
	_ = rows.Close()

	switch {
	case hasOld && hasNew:
		return fmt.Errorf("%s에 %s와 %s 컬럼이 동시에 존재합니다", table, oldColumn, newColumn)
	case !hasOld && !hasNew:
		return fmt.Errorf("%s에 %s 컬럼이 없습니다", table, oldColumn)
	case hasNew:
		return nil
	}
	if _, err := db.ExecContext(ctx,
		fmt.Sprintf("ALTER TABLE %s RENAME COLUMN %s TO %s", table, oldColumn, newColumn)); err != nil {
		return fmt.Errorf("%s.%s → %s 변경 실패: %w", table, oldColumn, newColumn, err)
	}
	return nil
}
