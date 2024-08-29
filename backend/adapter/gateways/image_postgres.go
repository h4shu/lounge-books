package gateways

import (
	"context"
	_ "embed"

	"github.com/h4shu/lounge-books/adapter/gateways/dto"
	"github.com/h4shu/lounge-books/application/repositories"
	"github.com/h4shu/lounge-books/domain/entities"
	"github.com/h4shu/lounge-books/domain/valueobjects"
	_ "github.com/lib/pq"
)

type ImagePostgres struct {
	db repositories.SQL
}

//go:embed schema/image_postgres.sql
var imagePostgresSchema string

func NewImagePostgres(db repositories.SQL) (repositories.ImageRepository, error) {
	err := db.Exec(imagePostgresSchema)
	if err != nil {
		return nil, err
	}
	return &ImagePostgres{
		db: db,
	}, nil
}

func (p *ImagePostgres) Create(ctx context.Context, image *entities.Image) error {
	imageDTO := dto.NewImageDTO(image)
	query := "INSERT INTO images(id, s3_key, format, original_url) VALUES($1, $2, $3, $4)"
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.ExecContext(ctx, imageDTO.ID(), imageDTO.S3Key(), imageDTO.Format(), imageDTO.OriginalURL())
}

func (p *ImagePostgres) FindByID(ctx context.Context, imageId *valueobjects.ImageID) (*entities.Image, error) {
	query := "SELECT id, s3_key, format, original_url FROM images WHERE id = $1"
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var (
		id          string
		isbn        string
		title       string
		s3Key       string
		format      string
		originalURL string
	)
}

func (p *ImagePostgres) FindByBookID(ctx context.Context, bookId *valueobjects.BookID) ([]entities.Image, error) {
}
