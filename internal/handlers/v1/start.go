package v1

import (
	"github.com/MrWebUzb/voovozbot/internal/models"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
)

func (h *HandlerV1) Start(m *telebot.Message) {
	h.log.Info("start command called")

	_ = h.strg.User().Upsert(&models.User{
		ID:        m.Sender.ID,
		Firstname: m.Sender.FirstName,
		Lastname:  m.Sender.LastName,
		Username:  m.Sender.Username,
	})

	msg, err := h.b.Send(m.Sender, "Hello there!")

	h.log.Info("sent data", zap.Any("msg", msg), zap.Error(err))
}
