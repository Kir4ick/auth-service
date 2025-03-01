package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/pkg/errors"
)

type Config struct {
	Network
	Database
	RabbitMq
	Mail
}

func New() (*Config, error) {
	config := &Config{}

	err := env.Parse(config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse env")
	}

	return config, nil
}

type Network struct {
	HTTP_PORT string `env:"HTTP_PORT"`
	GRPC_PORT string `env:"GRPC_PORT"`
}

type Database struct {
	HOST     string `env:"DB_HOST"`
	PORT     string `env:"DB_PORT"`
	USER     string `env:"DB_USER"`
	PASS     string `env:"DB_PASS"`
	DATABASE string `env:"DB_DATABASE"`
}

type RabbitMq struct {
	USER string `env:"RABBIT_USER"`
	PASS string `env:"RABBIT_PASS"`
	PORT string `env:"RABBIT_PORT"`
}

type Mail struct {
	MAILER       string `env:"MAIL_MAILER"`
	SCHEME       string `env:"MAIL_SCHEME"`
	HOST         string `env:"MAIL_HOST"`
	PORT         string `env:"MAIL_PORT"`
	USERNAME     string `env:"MAIL_USERNAME"`
	PASSWORD     string `env:"MAIL_PASSWORD"`
	FROM_ADDRESS string `env:"MAIL_FROM_ADDRESS"`
	FROM_NAME    string `env:"MAIL_FROM_NAME"`
}
