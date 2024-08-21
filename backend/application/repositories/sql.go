package repositories

import "context"

type SQL interface {
	Exec(query string, args ...any) error
	ExecContext(context.Context, string, ...any) error
	QueryContext(context.Context, string, ...any) (Rows, error)
	QueryRowContext(context.Context, string, ...any) Row
	PrepareContext(ctx context.Context, query string) (Stmt, error)
}

type Rows interface {
	Scan(dest ...any) error
	Next() bool
	Err() error
	Close() error
}

type Row interface {
	Scan(dest ...any) error
}

type Stmt interface {
	ExecContext(ctx context.Context, args ...any) error
	QueryRowContext(ctx context.Context, args ...any) Row
	Close() error
}
