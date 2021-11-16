package config

import "github.com/MrWebUzb/goenv"

type Config struct {
	AppEnvironment string `env:"APP_ENVIRONMET" default:"develop"`
	BotToken       string `env:"BOT_TOKEN"`
	ChannelID      int64  `env:"CHANNEL_ID"`
	// PostgresHost     string `env:"POSTGRES_HOST"`
	// PostgresPort     int32  `env:"POSTGRES_PORT"`
	// PostgresUser     string `env:"POSTGRES_USER"`
	// PostgresDatabase string `env:"POSTGRES_DATABASE"`
	// PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresURL string `env:"DATABASE_URL"`
}

func New(fileNames ...string) (*Config, error) {
	env, err := goenv.New(fileNames...)

	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := env.Parse(cfg, nil); err != nil {
		return nil, err
	}

	return cfg, nil
}
