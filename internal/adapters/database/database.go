package database

import (
	"context"

	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/entities"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type PgLinkStorage struct {
	db *sqlx.DB
}

func NewPgEventStorage(dsn string) (*PgLinkStorage, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgLinkStorage{db: db}, nil
}

func (pges *PgLinkStorage) GetLinkByFrom(from string) (string, error) {
	var toLink string

	query := `
		SELECT toLink
		FROM links
		WHERE fromLink = $1
	`

	err := pges.db.GetContext(context.Background(), &toLink, query, from)
	if err != nil {
		return "", err
	}

	return toLink, nil
}

func (pges *PgLinkStorage) SaveLink(link entities.Link) error {
	query := `
		INSERT INTO links (fromlink, tolink)
		VALUES ($1, $2);
	`
	_, err := pges.db.ExecContext(context.Background(), query, link.From, link.To)
	return err
}
