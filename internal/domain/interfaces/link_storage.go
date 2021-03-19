package interfaces

import (
	"github.com/alekseysychev/avito-auto-backend-trainee-assignment/internal/domain/entities"
)

type LinkStorage interface {
	GetLinkByFrom(from string) (string, error)
	SaveLink(link entities.Link) bool
}
