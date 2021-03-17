package usecases

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/entities"
	err "github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/errors"
	"github.com/stretchr/testify/assert"
)

type LinkStorageTest struct {
	s string
	e error
}

func (lst *LinkStorageTest) GetLinkByFrom(from string) (string, error) {
	return lst.s, lst.e
}

func (lst *LinkStorageTest) SaveLink(link entities.Link) error {
	return lst.e
}

func Test(t *testing.T) {
	t.Run("GetLinkByFrom", func(t *testing.T) {
		for _, test := range []struct {
			n string
			f string
			t string
			e error
		}{
			{"empty from", "", "", err.ErrEmptyFromLink},
			{"ot empty from", "/from", "/to", nil},
		} {
			service := LinkService{
				LinkStorage: &LinkStorageTest{
					s: test.t,
					e: test.e,
				},
			}
			to, err := service.GetLinkByFrom(test.f)
			assert.Equal(t, to, test.t, "should be equal "+test.n)
			assert.Equal(t, err, test.e, "should be equal "+test.n)
		}
	})
	t.Run("generateRandomLink len", func(t *testing.T) {
		for _, test := range []struct {
			na string
			n  int
			w  int
		}{
			{"zero", 0, 6},
			{"low", rand.Int() * -1, 6},
			{"max", rand.Int() + 21, 6},
			{"six", 6, 6},
			{"four", 4, 4},
		} {
			r := generateRandomLink(test.n)
			assert.Equal(t, len(r), test.w, "should be equal "+test.na)
		}
	})
	t.Run("generateRandomLink just generate", func(t *testing.T) {
		r1 := generateRandomLink(6)
		r2 := generateRandomLink(6)
		assert.NotEqual(t, r1, r2, "should be not equal")
	})
	t.Run("CreateLink", func(t *testing.T) {
		for _, test := range []struct {
			n string
			f string
			t string
			e error
		}{
			{"empty from & to", "", "", err.ErrEmptyToLink},
			{"not empty from & to", "/from", "/to", nil},
			{"not empty from & to But databaseError", "/from", "/to", errors.New("DB error")},
		} {
			service := LinkService{
				LinkStorage: &LinkStorageTest{
					s: test.f,
					e: test.e,
				},
			}
			err := service.CreateLink(entities.Link{
				From: test.f,
				To:   test.t,
			})
			assert.Equal(t, err, test.e, "should be equal"+test.n)
		}
	})
}
