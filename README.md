# SQLCore


```go
package sqlcore

type Executor interface {
    ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
    QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
    QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

func Int64(value any) int64
func ExecIndexes(ctx context.Context, db Executor, queries ...string) error
```

## 사용

```go
import "github.com/G226Medi03/SQLCore"
```

## 개발

```sh
go test ./...
```
