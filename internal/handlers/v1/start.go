package v1

import (
	"github.com/MrWebUzb/voovozbot/internal/models"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
)

func (h *HandlerV1) Start(m *telebot.Message) {
	h.log.Info("start command called")

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

	msg, err := h.b.Send(m.Sender, "Hello there!")

	h.log.Info("sent data", zap.Any("msg", msg), zap.Error(err))
}
