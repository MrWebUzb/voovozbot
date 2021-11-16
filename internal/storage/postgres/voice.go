package postgres

import (
	"github.com/MrWebUzb/voovozbot/internal/models"
	"github.com/MrWebUzb/voovozbot/internal/storage/repo"
	"github.com/jmoiron/sqlx"
)

type voiceRepo struct {
	db *sqlx.DB
}

var _ repo.VoiceI = &voiceRepo{}

func NewVoice(db *sqlx.DB) repo.VoiceI {
	return &voiceRepo{
		db: db,
	}
}

func (cr *voiceRepo) Upsert(req *models.Voice) error {
	query := `
		INSERT INTO "voices" (
			file_unique_id,
			file_id,
			duration,
			mime_type,
			file_size,
			caption
		) VALUES($1, $2, $3, $4, $5, $6)
		ON CONFLICT(file_unique_id) DO UPDATE SET
			file_id=$2,
			duration=$3,
			mime_type=$4,
			file_size=$5,
			caption=$6;
	`

	_, err := cr.db.Exec(query,
		req.FileUniqueID,
		req.FileID,
		req.Duration,
		req.MimeType,
		req.FileSize,
		req.Caption,
	)

	return err
}

func (cr *voiceRepo) Search(search string, offset, limit int) ([]*models.Voice, error) {
	query := `
		SELECT
			file_unique_id,
			file_id,
			duration,
			mime_type,
			file_size,
			caption
		FROM voices
	`

	where := ""
	offsetAndLimit := " OFFSET :offset LIMIT :limit"

	if search != "" {
		where = " WHERE caption ilike '%' || :search || '%'"
	}

	params := map[string]interface{}{
		"search": search,
		"offset": offset,
		"limit":  limit,
	}

	stmt, err := cr.db.PrepareNamed(query + where + offsetAndLimit)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resp []*models.Voice

	for rows.Next() {
		var voice models.Voice

		if err := rows.Scan(
			&voice.FileUniqueID,
			&voice.FileID,
			&voice.Duration,
			&voice.MimeType,
			&voice.FileSize,
			&voice.Caption,
		); err != nil {
			return nil, err
		}

		resp = append(resp, &voice)
	}

	return resp, nil
}
