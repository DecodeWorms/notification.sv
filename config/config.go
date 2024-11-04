package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	appEnv      = "APP_ENV"
	servicePort = "APP_PORT"
	password    = "SMTP_PASSWORD"
	pulsarUrl   = "PULSAR_URL"
)

type source interface {
	GetEnv(key string, fallback string) string
	GetEnvBool(key string, fallback bool) bool
	GetEnvInt(key string, fallback int) int
}

type OSSource struct {
	source //nolint
}

func (o OSSource) GetEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func (o OSSource) GetEnvBool(key string, fallback bool) bool {
	b := o.GetEnv(key, "")
	if len(b) == 0 {
		return fallback
	}
	v, err := strconv.ParseBool(b)
	if err != nil {
		return fallback
	}
	return v
}

func (o OSSource) GetEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		result, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return result
	}
	return fallback
}

type Config struct {
	Password    string
	AppEnv      string
	ServicePort string
	PulsarUrl   string
}

func ImportConfig(source source) Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appEnv := source.GetEnv(appEnv, "")
	port := source.GetEnv(servicePort, "8001")
	url := source.GetEnv(pulsarUrl, "")
	pass := source.GetEnv(password, "")

	return Config{
		AppEnv:      appEnv,
		ServicePort: port,
		PulsarUrl:   url,
		Password:    pass,
	}
}
