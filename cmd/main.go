package main

import (
	"github.com/MrWebUzb/voovozbot/internal/bot"
	"github.com/MrWebUzb/voovozbot/internal/config"
	"github.com/MrWebUzb/voovozbot/internal/storage"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

var logger *zap.Logger

func main() {
	cfg, err := config.New()

	if err != nil {
		panic(err)
	}

	// psqlUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	cfg.PostgresHost,
	// 	cfg.PostgresPort,
	// 	cfg.PostgresUser,
	// 	cfg.PostgresPassword,
	// 	cfg.PostgresDatabase,
	// )

	psqlConn, err := sqlx.Connect("postgres", cfg.PostgresURL)
	if err != nil {
		panic(err)
	}

	if cfg.AppEnvironment == "develop" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	strg := storage.NewStoragePg(psqlConn)

	bot, err := bot.NewBot(cfg, logger, strg)

	if err != nil {
		panic(err)
	}

	bot.Start()
}
