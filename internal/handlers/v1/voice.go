package v1

import (
	"strings"

	"github.com/MrWebUzb/voovozbot/internal/models"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
)

func (h *HandlerV1) OnVoiceSentToChannel(m *telebot.Message) {
	h.log.Info("message received", zap.Any("message", m))

	v := m.Voice

	if v == nil {
		return
	}

	caption := m.Caption

	if caption == "" {
		caption = v.Caption
	}

	caption = strings.ReplaceAll(caption, "@Ovoz_Parcha", "")

	caption = strings.TrimSpace(caption)

	voice := &models.Voice{
		Duration:     int64(v.Duration),
		MimeType:     v.MIME,
		FileID:       v.FileID,
		FileUniqueID: v.UniqueID,
		FileSize:     int64(v.FileSize),
		Caption:      caption,
	}

	if err := h.strg.Voice().Upsert(voice); err != nil {
		h.log.Error("could not save voice to database", zap.Error(err))
		return
	}

	h.log.Info("successfully saved", zap.Any("voice", voice))
}
