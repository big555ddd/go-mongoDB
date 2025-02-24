package config

import (
	"app/internal/logger"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Init() {
	godotenv.Load()
	app()
	logger.InitLogger()
	Database()
	OAuth()
}

func app() {
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("APP_ENV", "development")

	conf("APP_ENV", "development")
	conf("APP_PORT", "8080")
	conf("APP_NAME", "")
	conf("DEBUG", "false")

	conf("DB_HOST", "")
	conf("DB_PORT", "")
	conf("DB_USER", "")
	conf("DB_PASSWORD", "")
	conf("DB_DATABASE", "")
	conf("DB_SRV", "false")

	conf("TOKEN_SECRET_KEY", "secret")
	conf("TOKEN_DURATION", "24h")

	conf("EMAIL_HOST", "")
	conf("EMAIL_PORT", "")
	conf("EMAIL_USERNAME", "")
	conf("EMAIL_PASSWORD", "")
}
