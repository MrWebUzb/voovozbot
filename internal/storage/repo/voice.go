package repo

import "github.com/MrWebUzb/voovozbot/internal/models"

type VoiceI interface {
	Upsert(req *models.Voice) error
	Search(search string, offset, limit int) ([]*models.Voice, error)
}
