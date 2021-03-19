package database

// import (
// 	"testing"

// 	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/entities"
// 	sqlxmock "github.com/zhashkevych/go-sqlxmock"
// )

// type DBTest struct{}

// func NewPgEventStorageMock(t *testing.T) (*PgLinkStorage, sqlxmock.Sqlmock) {
// 	db, mock, err := sqlxmock.Newx()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	return &PgLinkStorage{db: db}, mock
// }

// func TestGetLinkByFrom(t *testing.T) {
// 	linkStorage, mock := NewPgEventStorageMock(t)

// 	tests := []struct {
// 		name    string
// 		s       *PgLinkStorage
// 		link    entities.Link
// 		mock    func()
// 		want    int
// 		wantErr bool
// 	}{
// 		{
// 			name: "OK",
// 			s:    linkStorage,
// 			link: entities.Link{
// 				From: "/from",
// 				To:   "/to",
// 			},
// 			mock: func() {
// 				// rows := sqlxmock.NewRows([]string{"id"}).AddRow(1)
// 				// rows := sqlxmock.NewRows([]string{"id"})

// 				// rows := sqlxmock.NewRows([]string{"fromlink", "tolink"}).AddRow("/from", "/to")
// 				mock.ExpectQuery("INSERT INTO links").WithArgs("fromlink", "tolink")
// 				//.WillReturnRows(rows)
// 			},
// 			want: 1,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mock()
// 			err := tt.s.SaveLink(tt.link)

// 			t.Error(err)
// 			// if (err != nil) != tt.wantErr {
// 			// t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
// 			// return
// 			// }
// 		})
// 	}
// }
