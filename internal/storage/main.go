package storage

import (
	"github.com/MrWebUzb/voovozbot/internal/storage/postgres"
	"github.com/MrWebUzb/voovozbot/internal/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Voice() repo.VoiceI
	User() repo.UserI
}

type storagePg struct {
	voiceRepo repo.VoiceI
	userRepo  repo.UserI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		voiceRepo: postgres.NewVoice(db),
		userRepo:  postgres.NewUser(db),
	}
}

func (s *storagePg) Voice() repo.VoiceI {
	return s.voiceRepo
}

func (s *storagePg) User() repo.UserI {
	return s.userRepo
}
