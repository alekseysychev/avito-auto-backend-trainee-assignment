package database

import (
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
	return check(db)
}

func check(db *sqlx.DB) (*PgLinkStorage, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgLinkStorage{db: db}, nil
}

func (pges *PgLinkStorage) GetLinkByFrom(from string) (string, error) {
	var toLink string

	row := pges.db.QueryRow("SELECT toLink FROM links WHERE fromLink = $1", from)

	if err := row.Scan(&toLink); err != nil {
		return "", err
	}

	return toLink, nil
}

func (pges *PgLinkStorage) SaveLink(link entities.Link) bool {
	row := pges.db.QueryRow("INSERT INTO links (fromlink, tolink) VALUES ($1, $2) ON CONFLICT DO NOTHING RETURNING true;",
		link.From, link.To)

	var success bool
	if err := row.Scan(&success); err != nil {
		return false
	}

	return success
}
