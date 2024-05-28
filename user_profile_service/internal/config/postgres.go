package config

import (
	"context"
	"log"
	"time"

	"github.com/hollyhox-21/discord_project/libraries/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

type Postgres struct {
	PgDSN               string        `env:"USER_PROFILE_POSTGRES_PG_DSN" env-required:"true" env-description:"postgres string connection"`
	MaxConnectionsCount int32         `env:"USER_PROFILE_POSTGRES_MAX_OPEN_CONN" env-default:"10" env-description:"postgres max count connection"`
	MinConnectionsCount int32         `env:"USER_PROFILE_POSTGRES_MIN_OPEN_CONN" env-default:"2" env-description:"postgres min count connection"`
	MaxConnIdleTime     time.Duration `env:"USER_PROFILE_POSTGRES_MAX_IDLE_TIME" env-default:"1h" env-description:"postgres max connection idle time"`
	MaxConnLifeTime     time.Duration `env:"USER_PROFILE_POSTGRES_MAX_LIFE_TIME" env-default:"1m" env-description:"postgres max connection lifetime"`
}

// Parse forms the postgres config
func (c *Postgres) Parse(ctx context.Context) {
	if checkEnvFileExists() {
		err := cleanenv.ReadConfig(".env", c)
		if err != nil {
			logger.Fatalw(ctx, "Error loading file. Check .env file",
				"error", err)
		}
		return
	}

	err := cleanenv.ReadEnv(c)
	if err != nil {
		logger.Fatalw(ctx, "Error read env",
			"error", err)
	}
}

func (c *Postgres) Help() string {
	header := "Postgres config"
	help, err := cleanenv.GetDescription(c, &header)
	if err != nil {
		log.Fatalln(err)
	}
	return help
}
