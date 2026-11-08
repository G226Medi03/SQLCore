package sqlcore

import "context"

// ExecIndexes runs schema statements in order.
func ExecIndexes(ctx context.Context, db Executor, queries ...string) error {
	for _, query := range queries {
		if _, err := db.ExecContext(ctx, query); err != nil {
			return err
		}
	}
	return nil
}
