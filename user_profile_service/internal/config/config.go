package config

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	server "github.com/hollyhox-21/discord_project/user_profile_service/internal/server_v2/config"
)

type helper interface {
	Help() string
}

// Config is common config
type Config struct {
	Environment string `env:"ENV" env-default:"local" env-description:"environment for app"`

	ServerConfig   server.Config
	PostgresConfig Postgres
}

// Parse forms the config
func Parse(ctx context.Context) *Config {
	var config Config

	config.ServerConfig.Parse(ctx)
	config.PostgresConfig.Parse(ctx)

	helpInfo(
		&config.ServerConfig,
		&config.PostgresConfig,
	)

	return &config
}

func helpInfo(h ...helper) {
	helpFlag := flag.Bool("help", false, "Показать справку")
	helpShort := flag.Bool("h", false, "Короткая форма флага help")

	flag.Parse()

	if *helpFlag || *helpShort {
		for _, item := range h {
			fmt.Println(item.Help(), "\n")
		}
		os.Exit(0)
	}
}

func checkEnvFileExists() bool {
	_, err := os.Stat(".env")
	return !errors.Is(err, os.ErrNotExist)
}
