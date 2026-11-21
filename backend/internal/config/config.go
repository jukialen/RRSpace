package config

import (
	"os"
	"github.com/appwrite/sdk-for-go/v7/appwrite"
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
		AppEnv:            os.Getenv("APP_ENV"),
		Origin:            os.Getenv("ORIGIN"),
	BcPort:              os.Getenv("BC_PORT"),
		DBHost:            os.Getenv("DB_HOST"),
		DBPort:            os.Getenv("DB_PORT"),
		DBUser:            os.Getenv("DB_USER"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		DBName:            os.Getenv("DB_NAME"),
		RedisHost:         os.Getenv("REDIS_HOST"),
		RedisPort:         os.Getenv("REDIS_PORT"),
		R2Endpoint:        os.Getenv("R2_ENDPOINT"),
		R2Bucket:          os.Getenv("R2_BUCKET"),
		R2AccessKey:       os.Getenv("R2_ACCESS_KEY"),
		R2SecretKey:       os.Getenv("R2_SECRET_KEY"),
		R2AccountID:	     os.Getenv("R2_ACCOUNT_ID"),
		JWTSecret:    	   os.Getenv("JWT_SECRET"),
		AppwriteRegion:    os.Getenv("APPWRITE_REGION"),
		AppwriteProjectID: os.Getenv("APPWRITE_PROJECT_ID"),
		AppwriteAPIKey:    os.Getenv("APPWRITE_API_KEY"),
	}

	return cfg, nil
}

func os.Getenv(key, fallback string) string {
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

func configAppwrite() {
 appwrite.NewClient(
	appwrite.WithEndpoint("https://" + c.AppwriteRegion + ".cloud.appwrite.io/v1"),
	appwrite.WithProject(c.AppwriteProjectID),
	appwrite.WithKey(c.AppwriteAPIKey),
)	
}