package postgres

import (
	"github.com/MrWebUzb/voovozbot/internal/models"
	"github.com/MrWebUzb/voovozbot/internal/storage/repo"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

var _ repo.UserI = &userRepo{}

func NewUser(db *sqlx.DB) repo.UserI {
	return &userRepo{
		db: db,
	}
}

func (cr *userRepo) Upsert(req *models.User) error {
	query := `
		INSERT INTO "users" (
			id,
			first_name,
			last_name,
			username
		) VALUES($1, $2, $3, $4)
		ON CONFLICT(id) DO UPDATE SET
			first_name=$2,
			last_name=$3,
			username=$4
		;
	`

	_, err := cr.db.Exec(query,
		req.ID,
		req.Firstname,
		req.Lastname,
		req.Username,
	)

	return err
}
