package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	AppEnv string

	// Server
	Port string

	// PostgreSQL / Neon
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// Redis
	RedisHost string
	RedisPort string

	// Appwrite
	AppwriteEndpoint         string
	AppwriteProjectID        string
	AppwriteAPIKey           string
	AppwriteDatabaseID       string
	AppwriteUsersCollection  string

	// Cloudflare R2
	R2Endpoint   string
	R2Bucket     string
	R2AccessKey  string
	R2SecretKey  string
	R2AccountID  string

	// Auth
	JWTSecret string
}

func Load() (*Config, error) {
	cfg := &Config{
		AppEnv:            envOrDefault("APP_ENV", "development"),
		Port:              envOrDefault("PORT", "8080"),
		DBHost:            envOrDefault("DB_HOST", "localhost"),
		DBPort:            envOrDefault("DB_PORT", "5432"),
		DBUser:            envOrDefault("DB_USER", "postgres"),
		DBPassword:        envOrDefault("DB_PASSWORD", "postgres"),
		DBName:            envOrDefault("DB_NAME", "studio_booking"),
		RedisHost:         envOrDefault("REDIS_HOST", "localhost"),
		RedisPort:         envOrDefault("REDIS_PORT", "6379"),
		R2Endpoint:      envOrDefault("R2_ENDPOINT", ""),
		R2Bucket:        envOrDefault("R2_BUCKET", "studio-photos"),
		R2AccessKey:     os.Getenv("R2_ACCESS_KEY"),
		R2SecretKey:     os.Getenv("R2_SECRET_KEY"),
		R2AccountID:     os.Getenv("R2_ACCOUNT_ID"),
		JWTSecret:       envOrDefault("JWT_SECRET", "dev-secret"),
	}

	return cfg, nil
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func (c *Config) DBDSN() string {
	return "host=" + c.DBHost +
		" port=" + c.DBPort +
		" user=" + c.DBUser +
		" password=" + c.DBPassword +
		" dbname=" + c.DBName +
		" sslmode=require"
}

func (c *Config) RedisAddr() string {
	return c.RedisHost + ":" + c.RedisPort
}