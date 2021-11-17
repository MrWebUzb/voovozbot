package v1

import (
	"strings"

	"github.com/MrWebUzb/voovozbot/internal/models"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
)

func (h *HandlerV1) OnVoiceSentToChannel(m *telebot.Message) {
	h.log.Info("message received", zap.Any("message", m))

	var user models.User

	if m.Sender == nil {
		user = models.User{
			ID:        int(m.Chat.ID),
			Firstname: m.Chat.FirstName,
			Lastname:  m.Chat.LastName,
			Username:  m.Chat.Username,
		}
	} else {
		user = models.User{
			ID:        m.Sender.ID,
			Firstname: m.Sender.FirstName,
			Lastname:  m.Sender.LastName,
			Username:  m.Sender.Username,
		}
	}
	_ = h.strg.User().Upsert(&user)

	if m.Chat != nil && m.Chat.ID != h.channelID {
		h.log.Warn("another channel used bot")
		return
	}

	v := m.Voice

	if v == nil {
		return
	}

	caption := m.Caption

	if caption == "" {
		caption = v.Caption
	}

	idx := strings.Index(caption, "\n")

	if idx > 0 {
		caption = strings.TrimSpace(caption[:idx])
	}

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

func (h *HandlerV1) OnVoiceChosen(q *telebot.ChosenInlineResult) {
	h.log.Info("chosen voice handler", zap.Any("query", q))
	_ = h.strg.User().Upsert(&models.User{
		ID:        q.From.ID,
		Firstname: q.From.FirstName,
		Lastname:  q.From.LastName,
		Username:  q.From.Username,
	})

	if err := h.strg.Voice().IncrementUsageCount(q.ResultID); err != nil {
		h.log.Error("could not update voice usage", zap.Error(err))
	}
}
