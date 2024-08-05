package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"

	_ "github.com/lib/pq"
)

type Config struct {
	GIN_SERVER_PORT    string
	GRPC_SERVER_PORT   string
	SECRET_KEY_ACCESS  string
	SECRET_KEY_REFRESH string
	HOST               string
	DB_PORT            string
	DB_HOST            string
	DB_USER            string
	DB_PASSWORD        string
	DB_NAME            string
	SMTP_AUTH_PASSWORD string
	SMTP_PORT          string
	SMTP_SERVER        string
	SMTP_SENDER        string
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	config := Config{}

	config.GIN_SERVER_PORT = cast.ToString(coalesce("GIN_SERVER_PORT", ":8081"))
	config.GRPC_SERVER_PORT = cast.ToString(coalesce("GRPC_SERVER_PORT", ":50050"))
	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "ecommerce_auth_service"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "123321"))
	config.SECRET_KEY_ACCESS = cast.ToString(coalesce("SECRET_KEY_ACCESS", "secret_key"))
	config.SECRET_KEY_REFRESH = cast.ToString(coalesce("SECRET_KEY_REFRESH", "not so easy"))
	config.SMTP_AUTH_PASSWORD = cast.ToString(coalesce("SMTP_AUTH_PASSWORD", "your_smtp_password"))
	config.SMTP_PORT = cast.ToString(coalesce("SMTP_PORT", "443"))
	config.SMTP_SERVER = cast.ToString(coalesce("SMTP_SERVER", "smtp.gmail.com"))
	config.SMTP_SENDER = cast.ToString(coalesce("SMTP_SENDER", "your_email@gmail.com"))

	return config
}

func coalesce(env string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(env)
	if !exists {
		return defaultValue
	}
	return value
}
