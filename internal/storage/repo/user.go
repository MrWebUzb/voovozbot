package repo

import "github.com/MrWebUzb/voovozbot/internal/models"

type UserI interface {
	Upsert(req *models.User) error
}
