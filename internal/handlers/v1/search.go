package v1

import (
	"strconv"

	"github.com/MrWebUzb/voovozbot/internal/models"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
)

func (h *HandlerV1) OnInlineSearch(q *telebot.Query) {
	h.log.Info("search query handled", zap.Any("query", q))

	_ = h.strg.User().Upsert(&models.User{
		ID:        q.From.ID,
		Firstname: q.From.FirstName,
		Lastname:  q.From.LastName,
		Username:  q.From.Username,
	})

	offset := parseOffset(q.Offset)
	limit := 10

	voices, err := h.strg.Voice().Search(q.Text, offset, limit)

	if err != nil {
		h.log.Error("error sending answer", zap.Error(err))
		h.EmptyAnswer(q)
		return
	}

	results := []telebot.Result{}

	for _, voice := range voices {
		var voiceRes telebot.VoiceResult

		voiceRes.ID = voice.FileUniqueID
		voiceRes.Title = voice.Caption
		voiceRes.Cache = voice.FileID
		voiceRes.Duration = int(voice.Duration)

		results = append(results, &voiceRes)
	}

	nextOffset := strconv.Itoa(offset + limit)

	if len(voices) == 0 {
		nextOffset = ""
	}

	if err := h.b.Answer(q, &telebot.QueryResponse{
		QueryID:    q.ID,
		Results:    results,
		NextOffset: nextOffset,
	}); err != nil {
		h.log.Error("error sending answer", zap.Error(err))
		h.EmptyAnswer(q)
		return
	}
}

func (h *HandlerV1) EmptyAnswer(q *telebot.Query) {
	h.b.Answer(q, &telebot.QueryResponse{
		QueryID: q.ID,
	})
}
