package gateways

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/repositories"
	"github.com/h4shu/lounge-books/domain/entities"
	_ "github.com/lib/pq"
)

type TagPostgres struct {
	db repositories.SQL
}

//go:embed schema/tag_postgres.sql
var schemaTagPostgres string

func NewTagPostgres(db repositories.SQL) (repositories.TagRepository, error) {
	err := db.Exec(schemaTagPostgres)
	if err != nil {
		return nil, err
	}
	return &TagPostgres{
		db: db,
	}, nil
}

func (p *TagPostgres) Create(ctx context.Context, tag *inputs.CreateTagInput) error {
	query := "INSERT INTO tags(name) VALUES($1)"
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.ExecContext(ctx, tag.Name)
}

func (p *TagPostgres) FindAll(ctx context.Context) ([]entities.Tag, error) {
	query := "SELECT id, name FROM tags"
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []entities.Tag
	for rows.Next() {
		var (
			id   int
			name string
		)
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}

		tag := entities.NewTag(id, name)
		tags = append(tags, *tag)
	}
	return tags, nil
}

func (p *TagPostgres) FindByID(ctx context.Context, tagId int) (*entities.Tag, error) {
	query := "SELECT id, name FROM tags WHERE id = $1"
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var (
		id   int
		name string
	)
	err = stmt.QueryRowContext(ctx, tagId).Scan(&id, &name)
	if err == sql.ErrNoRows {
		return nil, entities.ErrTagNotFound
	} else if err != nil {
		return nil, err
	}

	return entities.NewTag(id, name), nil
}
