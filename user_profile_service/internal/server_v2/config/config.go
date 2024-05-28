package config

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/hollyhox-21/discord_project/libraries/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

// Config - server config
type Config struct {
	GRPCServer        GRPCServer
	GRPCGatewayServer GRPCGatewayServer
	INFOServer        INFOServer
}

type GRPCServer struct {
	Addr string `env:"SERVER_GRPC_ADDR" env-default:":8081" env-description:"gRPC server port"`
}

type GRPCGatewayServer struct {
	Addr string `env:"SERVER_HTTP_PUBLIC_ADDR" env-default:":8080" env-description:"HTTP public server port"`
}

type INFOServer struct {
	Addr string `env:"SERVER_HTTP_INFO_ADDR" env-default:":8082" env-description:"HTTP info server port"`
}

// Parse forms the server config
func (c *Config) Parse(ctx context.Context) {
	if checkEnvFileExists() {
		err := cleanenv.ReadConfig(".env", c)
		if err != nil {
			logger.Fatalf(ctx, "Error loading .env file: %v", err)
		}
	}
	err := cleanenv.ReadEnv(c)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Config) Help() string {
	header := "Server config"
	help, err := cleanenv.GetDescription(c, &header)
	if err != nil {
		log.Fatalln(err)
	}
	return help
}

func checkEnvFileExists() bool {
	_, err := os.Stat(".env")
	return !errors.Is(err, os.ErrNotExist)
}
