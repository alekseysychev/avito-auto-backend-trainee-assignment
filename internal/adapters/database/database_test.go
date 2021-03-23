package database

import (
	"testing"

	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

type DBTest struct{}

func NewPgEventStorageMock(t *testing.T) (*PgLinkStorage, sqlxmock.Sqlmock) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	obj, err := check(db)
	if err != nil {
		t.Fatalf("an error '%s' was in check", err)
	}

	return obj, mock
}

func Test_GetLinkByFrom(t *testing.T) {
	linkStorage, mock := NewPgEventStorageMock(t)

	tests := []struct {
		name     string
		mock     func()
		linkFrom string
		linkTo   string
		wantErr  bool
	}{
		{
			name: "Found",
			mock: func() {
				rows := sqlxmock.NewRows([]string{"fromlink"}).AddRow("/to")
				mock.ExpectQuery("SELECT toLink FROM links").WithArgs("/from").WillReturnRows(rows)
			},
			linkFrom: "/from",
			linkTo:   "/to",
			wantErr:  false,
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlxmock.NewRows([]string{"fromlink"})
				mock.ExpectQuery("SELECT toLink FROM links").WithArgs("/from").WillReturnRows(rows)
			},
			linkFrom: "/from",
			linkTo:   "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			to, err := linkStorage.GetLinkByFrom("/from")
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.linkTo, to)
			}
		})
	}
}

func Test_SaveLink(t *testing.T) {
	linkStorage, mock := NewPgEventStorageMock(t)

	tests := []struct {
		name    string
		mock    func()
		link    entities.Link
		success bool
		wantErr bool
	}{
		{
			name: "Success",
			mock: func() {
				rows := sqlxmock.NewRows([]string{"success"}).AddRow("true")
				mock.ExpectQuery("INSERT INTO links").WithArgs("/from", "/to").WillReturnRows(rows)
			},
			link:    entities.Link{From: "/from", To: "/to"},
			success: true,
		},
		{
			name: "Error",
			mock: func() {
				rows := sqlxmock.NewRows([]string{"success"}).AddRow("false")
				mock.ExpectQuery("INSERT INTO links").WithArgs("/from", "/to").WillReturnRows(rows)
			},
			link:    entities.Link{From: "/from", To: "/true"},
			success: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			success := linkStorage.SaveLink(tt.link)
			// fmt.Println(mock.ExpectationsWereMet())
			assert.Equal(t, tt.success, success, "Test name: %s", tt.name)
		})
	}
}
