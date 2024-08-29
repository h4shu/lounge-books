package gateways

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/h4shu/lounge-books/application/repositories"
	"github.com/h4shu/lounge-books/domain/entities"
	"github.com/h4shu/lounge-books/domain/valueobjects"
	_ "github.com/lib/pq"
)

type BookPostgres struct {
	db repositories.SQL
}

//go:embed schema/book_postgres.sql
var schema string

func NewBookPostgres(db repositories.SQL) (repositories.BookRepository, error) {
	err := db.Exec(schema)
	if err != nil {
		return nil, err
	}
	return &BookPostgres{
		db: db,
	}, nil
}

func (p *BookPostgres) Create(ctx context.Context, book *entities.Book) error {
	query := "INSERT INTO books(isbn, title, description, cover_link, published_year, published_month, published_day, author, publisher, page_count) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.ExecContext(ctx, book.ISBN().String(), book.Title(), book.Description(), book.CoverLink(), book.PublishedAt().Year().Int(), book.PublishedAt().Month().Int(), book.PublishedAt().Day().Int(), book.Author().String(), book.Publisher(), book.PageCount())
}

func (p *BookPostgres) FindAll(ctx context.Context) ([]entities.Book, error) {
	query := "SELECT id, isbn, title, description, cover_link, published_year, published_month, published_day, author, publisher, page_count, deleted_at FROM books"
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entities.Book
	for rows.Next() {
		var (
			id             int
			isbn           string
			title          string
			description    string
			coverLink      string
			publishedYear  int
			publishedMonth int
			publishedDay   int
			author         string
			publisher      string
			pageCount      int
			deletedAt      sql.NullTime
		)
		err := rows.Scan(&id, &isbn, &title, &description, &coverLink, &publishedYear, &publishedMonth, &publishedDay, &author, &publisher, &pageCount, &deletedAt)
		if err != nil {
			return nil, err
		}

		publishedAt, err := valueobjects.NewPublishedAt(publishedYear, publishedMonth, publishedDay)
		if err != nil {
			return nil, err
		}

		var deletedAtVal *valueobjects.DeletedAt
		if deletedAt.Valid {
			deletedAtVal = valueobjects.NewDeletedAt(&deletedAt.Time)
		} else {
			deletedAtVal = valueobjects.NewDeletedAt(nil)
		}

		book := entities.NewBook(valueobjects.NewBookID(id), valueobjects.NewISBN(isbn), title, description, coverLink, publishedAt, valueobjects.NewAuthor(author), publisher, pageCount, deletedAtVal)
		books = append(books, *book)
	}
	return books, nil
}

func (p *BookPostgres) FindByID(ctx context.Context, bookId *valueobjects.BookID) (*entities.Book, error) {
	query := "SELECT id, isbn, title, description, cover_link, published_year, published_month, published_day, author, publisher, page_count, deleted_at FROM books WHERE id = $1"
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var (
		id             int
		isbn           string
		title          string
		description    string
		coverLink      string
		publishedYear  int
		publishedMonth int
		publishedDay   int
		author         string
		publisher      string
		pageCount      int
		deletedAt      sql.NullTime
	)
	err = stmt.QueryRowContext(ctx, bookId.Int()).Scan(&id, &isbn, &title, &description, &coverLink, &publishedYear, &publishedMonth, &publishedDay, &author, &publisher, &pageCount, &deletedAt)
	if err == sql.ErrNoRows {
		return nil, entities.ErrBookNotFound
	} else if err != nil {
		return nil, err
	}

	publishedAt, err := valueobjects.NewPublishedAt(publishedYear, publishedMonth, publishedDay)
	if err != nil {
		return nil, err
	}

	var deletedAtVal *valueobjects.DeletedAt
	if deletedAt.Valid {
		deletedAtVal = valueobjects.NewDeletedAt(&deletedAt.Time)
	} else {
		deletedAtVal = valueobjects.NewDeletedAt(nil)
	}

	return entities.NewBook(valueobjects.NewBookID(id), valueobjects.NewISBN(isbn), title, description, coverLink, publishedAt, valueobjects.NewAuthor(author), publisher, pageCount, deletedAtVal), nil
}

func (p *BookPostgres) FindByKeywordContaining(ctx context.Context, keyword string) ([]entities.Book, error) {
	query := "SELECT id, isbn, title, description, cover_link, published_year, published_month, published_day, author, publisher, page_count, deleted_at FROM books WHERE isbn LIKE '%' || $1 || '%' OR lower(title) LIKE lower('%' || $1 || '%') OR lower(author) LIKE lower('%' || $1 || '%') OR lower(publisher) LIKE lower('%' || $1 || '%')"
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entities.Book
	for rows.Next() {
		var (
			id             int
			isbn           string
			title          string
			description    string
			coverLink      string
			publishedYear  int
			publishedMonth int
			publishedDay   int
			author         string
			publisher      string
			pageCount      int
			deletedAt      sql.NullTime
		)
		err := rows.Scan(&id, &isbn, &title, &description, &coverLink, &publishedYear, &publishedMonth, &publishedDay, &author, &publisher, &pageCount, &deletedAt)
		if err != nil {
			return nil, err
		}

		publishedAt, err := valueobjects.NewPublishedAt(publishedYear, publishedMonth, publishedDay)
		if err != nil {
			return nil, err
		}

		var deletedAtVal *valueobjects.DeletedAt
		if deletedAt.Valid {
			deletedAtVal = valueobjects.NewDeletedAt(&deletedAt.Time)
		} else {
			deletedAtVal = valueobjects.NewDeletedAt(nil)
		}

		book := entities.NewBook(valueobjects.NewBookID(id), valueobjects.NewISBN(isbn), title, description, coverLink, publishedAt, valueobjects.NewAuthor(author), publisher, pageCount, deletedAtVal)
		books = append(books, *book)
	}
	return books, nil
}

func (p *BookPostgres) Update(ctx context.Context, book *entities.Book) error {
	query := "UPDATE books SET isbn = $1, title = $2, description = $3, cover_link = $4, published_year = $5, published_month = $6, published_day = $7, author = $8, publisher = $9, page_count = $10, deleted_at = $11 WHERE id = $12"
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.ExecContext(ctx, book.ISBN().String(), book.Title(), book.Description(), book.CoverLink(), book.PublishedAt().Year().Int(), book.PublishedAt().Month().Int(), book.PublishedAt().Day().Int(), book.Author().String(), book.Publisher(), book.PageCount(), book.DeletedAt().String(), book.ID().Int())
}
