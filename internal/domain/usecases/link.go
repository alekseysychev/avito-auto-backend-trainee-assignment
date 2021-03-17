package usecases

import (
	"math/rand"

	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/entities"
	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/errors"
	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/interfaces"
)

type LinkService struct {
	LinkStorage interfaces.LinkStorage
}

func (ls *LinkService) GetLinkByFrom(from string) (string, error) {
	if from == "" {
		return "", errors.ErrEmptyFromLink
	}
	link, err := ls.LinkStorage.GetLinkByFrom(from)
	return link, err
}

func generateRandomLink(n int) string {
	if n <= 0 {
		n = 6
	}
	if n > 20 {
		n = 6
	}
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (ls *LinkService) CreateLink(link entities.Link) error {
	if link.From == "" {
		link.From = generateRandomLink(6)
	}
	if link.To == "" {
		return errors.ErrEmptyToLink
	}
	err := ls.LinkStorage.SaveLink(link)
	return err
}
