package api

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/entities"

	err "github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/errors"
	"github.com/stretchr/testify/assert"
)

type LinkUseCasesTest struct {
	name string
	f    string
	t    string
	e    error
}

func (luct *LinkUseCasesTest) GetLinkByFrom(from string) (string, error) {
	return luct.t, luct.e
}

func (luct *LinkUseCasesTest) CreateLink(requestData entities.Link) error {
	return luct.e
}

func Test(t *testing.T) {
	t.Run("GetLinkByFrom", func(t *testing.T) {
		for _, test := range []LinkUseCasesTest{
			{"empty from", "", "", err.ErrEmptyFromLink},
			{"not empty from", "/from", "/to", nil},
			{"usecase error", "/from", "/to", err.ErrEmptyToLink},
		} {
			server := &HttpServer{
				LinkUseCases: &test,
			}
			to, err := server.GetLinkByFrom(test.f)
			assert.Equal(t, to, test.t, "should be equal "+test.name)
			assert.Equal(t, err, test.e, "should be equal "+test.name)
		}
	})
	t.Run("CreateLink", func(t *testing.T) {
		for _, test := range []LinkUseCasesTest{
			{"error", "", "", err.ErrEmptyFromLink},
			{"not error", "/from", "/to", nil},
		} {
			server := &HttpServer{
				LinkUseCases: &test,
			}
			err := server.CreateLink(entities.Link{
				From: test.f,
				To:   test.t,
			})
			assert.Equal(t, err, test.e, "should be equal "+test.name)
		}
	})

	t.Run("Like integration POST", func(t *testing.T) {
		{
			req := httptest.NewRequest("POST", "/", nil)
			recorder := httptest.NewRecorder()
			server := &HttpServer{
				LinkUseCases: &LinkUseCasesTest{
					"not error", "/from", "/to", nil,
				},
			}
			server.ServeHTTP(recorder, req)
		}
		{
			jsonStr := []byte(`{"from":"/from", "to":"/to"}`)
			req := httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
			recorder := httptest.NewRecorder()
			server := &HttpServer{
				LinkUseCases: &LinkUseCasesTest{
					"not error", "/from", "/to", nil,
				},
			}
			server.ServeHTTP(recorder, req)
		}
		{
			jsonStr := []byte(`{"from":"", "to":""}`)
			req := httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
			recorder := httptest.NewRecorder()
			server := &HttpServer{
				LinkUseCases: &LinkUseCasesTest{
					"error", "", "", err.ErrEmptyFromLink,
				},
			}
			server.ServeHTTP(recorder, req)
		}
	})

	t.Run("Like integration GET", func(t *testing.T) {
		{
			req := httptest.NewRequest("GET", "/", nil)
			recorder := httptest.NewRecorder()
			server := &HttpServer{
				LinkUseCases: &LinkUseCasesTest{"error", "", "", err.ErrEmptyFromLink},
			}
			server.ServeHTTP(recorder, req)
		}
		{
			req := httptest.NewRequest("GET", "/", nil)
			recorder := httptest.NewRecorder()
			server := &HttpServer{
				LinkUseCases: &LinkUseCasesTest{"not error", "/from", "/to", nil},
			}
			server.ServeHTTP(recorder, req)
		}
	})
}
