// Package sqlcore provides the minimal SQL execution surface shared by the
// MedicalEntity repositories.
package sqlcore

import (
	"context"
	"database/sql"
)

// Executor is the minimal database surface used by the repositories.
// *sql.DB, *sql.Tx and the application store types satisfy it structurally.
type Executor interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}
