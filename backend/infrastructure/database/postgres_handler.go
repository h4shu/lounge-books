package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/h4shu/lounge-books/application/repositories"
)

type postgresHandler struct {
	db *sql.DB
}

func newPostgresHandler(c *config) (*postgresHandler, error) {
	ds := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.host,
		c.port,
		c.user,
		c.password,
		c.database,
	)
	db, err := sql.Open(c.driver, ds)
	if err != nil {
		return nil, err
	}
	return &postgresHandler{
		db: db,
	}, nil
}

func (p *postgresHandler) Exec(query string, args ...any) error {
	_, err := p.db.Exec(query, args...)
	return err
}

func (p *postgresHandler) ExecContext(ctx context.Context, query string, args ...any) error {
	_, err := p.db.ExecContext(ctx, query, args...)
	return err
}

func (p *postgresHandler) QueryContext(ctx context.Context, query string, args ...any) (repositories.Rows, error) {
	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return newPostgresRows(rows), nil
}

func (p *postgresHandler) QueryRowContext(ctx context.Context, query string, args ...any) repositories.Row {
	row := p.db.QueryRowContext(ctx, query, args...)
	return newPostgresRow(row)
}

func (p *postgresHandler) PrepareContext(ctx context.Context, query string) (repositories.Stmt, error) {
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return newPostgresStmt(stmt), nil
}

type postgresRow struct {
	row *sql.Row
}

func newPostgresRow(row *sql.Row) postgresRow {
	return postgresRow{row: row}
}

func (pr postgresRow) Scan(dest ...any) error {
	return pr.row.Scan(dest...)
}

type postgresRows struct {
	rows *sql.Rows
}

func newPostgresRows(rows *sql.Rows) postgresRows {
	return postgresRows{rows: rows}
}

func (pr postgresRows) Scan(dest ...any) error {
	return pr.rows.Scan(dest...)
}

func (pr postgresRows) Next() bool {
	return pr.rows.Next()
}

func (pr postgresRows) Err() error {
	return pr.rows.Err()
}

func (pr postgresRows) Close() error {
	return pr.rows.Close()
}

type postgresStmt struct {
	stmt *sql.Stmt
}

func newPostgresStmt(stmt *sql.Stmt) postgresStmt {
	return postgresStmt{stmt: stmt}
}

func (ps postgresStmt) ExecContext(ctx context.Context, args ...any) error {
	_, err := ps.stmt.ExecContext(ctx, args...)
	return err
}

func (ps postgresStmt) QueryRowContext(ctx context.Context, args ...any) repositories.Row {
	return ps.stmt.QueryRowContext(ctx, args...)
}

func (ps postgresStmt) Close() error {
	return ps.stmt.Close()
}
