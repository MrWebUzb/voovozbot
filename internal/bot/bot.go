package bot

import (
	"time"

	"github.com/MrWebUzb/voovozbot/internal/config"
	v1 "github.com/MrWebUzb/voovozbot/internal/handlers/v1"
	"github.com/MrWebUzb/voovozbot/internal/storage"
	"go.uber.org/zap"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Bot struct {
	b         *tb.Bot
	log       *zap.Logger
	strg      storage.StorageI
	channelID int64
}

func NewBot(cfg *config.Config, logger *zap.Logger, strg storage.StorageI) (*Bot, error) {
	settings := tb.Settings{
		Token:  cfg.BotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tb.NewBot(settings)

	if err != nil {
		return nil, err
	}

	return &Bot{
		b:         b,
		log:       logger,
		strg:      strg,
		channelID: cfg.ChannelID,
	}, nil
}

func (b *Bot) registerHandlers() {
	b.log.Info("registering handlers")
	handlerV1 := v1.NewHandlerV1(b.b, b.log, b.strg, b.channelID)

	b.b.Handle(StartCommand, handlerV1.Start)
	b.b.Handle(tb.OnChannelPost, handlerV1.OnVoiceSentToChannel)
	b.b.Handle(tb.OnEditedChannelPost, handlerV1.OnVoiceSentToChannel)

	b.b.Handle(tb.OnQuery, handlerV1.OnInlineSearch)

	b.b.Handle(tb.OnChosenInlineResult, handlerV1.OnVoiceChosen)
}

func (b *Bot) Start() {
	b.log.Info("bot started")
	b.registerHandlers()
	b.b.Start()
}
