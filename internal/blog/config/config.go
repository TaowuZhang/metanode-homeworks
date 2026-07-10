package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv         string
	ServerPort     string
	DBDriver       string
	DBDSN          string
	JWTSecret      string
	JWTExpireHours int
}

func Load() Config {
	_ = godotenv.Load()
	expire, err := strconv.Atoi(getenv("JWT_EXPIRE_HOURS", "24"))
	if err != nil || expire <= 0 {
		expire = 24
	}
	return Config{
		AppEnv:         getenv("APP_ENV", "development"),
		ServerPort:     getenv("SERVER_PORT", "8080"),
		DBDriver:       getenv("DB_DRIVER", "sqlite"),
		DBDSN:          getenv("DB_DSN", "blog.db"),
		JWTSecret:      getenv("JWT_SECRET", "change-me-in-local-env"),
		JWTExpireHours: expire,
	}
}

func (c Config) JWTDuration() time.Duration {
	return time.Duration(c.JWTExpireHours) * time.Hour
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
