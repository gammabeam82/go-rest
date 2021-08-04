package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	databaseUrl string
	httpPort    int
	jwtTokenTTL int
	jwtSecret   []byte
}

func (c Config) DatabaseUrl() string {
	return c.databaseUrl
}

func (c Config) HttpPort() int {
	return c.httpPort
}

func (c Config) JwtTokenTTL() int64 {
	return int64(c.jwtTokenTTL)
}

func (c Config) JwtSecret() []byte {
	return c.jwtSecret
}

func NewConfig() (*Config, error) {
	filename := ".env"
	env := os.Getenv("APP_ENV")

	switch env {
	case "", "dev":
		filename += ".local"
	case "test":
		filename += ".test"
	}

	err := godotenv.Load(filename)

	if err != nil {
		return nil, err
	}

	httpPort, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	tokenTTL, err := strconv.Atoi(os.Getenv("JWT_TOKEN_TTL"))

	if err != nil {
		return nil, err
	}

	return &Config{
		databaseUrl: os.Getenv("DATABASE_URL"),
		httpPort:    httpPort,
		jwtTokenTTL: tokenTTL,
		jwtSecret:   []byte(os.Getenv("JWT_SECRET")),
	}, nil
}
