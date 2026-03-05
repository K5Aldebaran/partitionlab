package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	ServiceHost    string
	ServicePort    int
	MinIOHost      string
	MinIOPort      string
	MinIOAccessKey string
	MinIOSecretKey string
	MinIOBucket    string
	MinIOUseSSL    bool
	// Redis для сессий
	RedisHost     string
	RedisPort     int
	RedisPassword string
	RedisDB       int
	// JWT
	JWTSecret string
}

func NewConfig() (*Config, error) {
	var err error

	configName := "config"
	_ = godotenv.Load()
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.WatchConfig()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	if host := os.Getenv("SERVICE_HOST"); host != "" {
		cfg.ServiceHost = host
	}
	if port := os.Getenv("SERVICE_PORT"); port != "" {
		if parsed, parseErr := strconv.Atoi(port); parseErr == nil {
			cfg.ServicePort = parsed
		}
	}

	if host := os.Getenv("REDIS_HOST"); host != "" {
		cfg.RedisHost = host
	}
	if port := os.Getenv("REDIS_PORT"); port != "" {
		if parsed, parseErr := strconv.Atoi(port); parseErr == nil {
			cfg.RedisPort = parsed
		}
	}
	if password := os.Getenv("REDIS_PASSWORD"); password != "" {
		cfg.RedisPassword = password
	}
	if db := os.Getenv("REDIS_DB"); db != "" {
		if parsed, parseErr := strconv.Atoi(db); parseErr == nil {
			cfg.RedisDB = parsed
		}
	}

	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		cfg.JWTSecret = jwtSecret
	}

	// MinIO configuration from environment
	cfg.MinIOHost = os.Getenv("MINIO_HOST")
	if cfg.MinIOHost == "" {
		cfg.MinIOHost = "127.0.0.1"
	}
	cfg.MinIOPort = os.Getenv("MINIO_PORT")
	if cfg.MinIOPort == "" {
		cfg.MinIOPort = "9000"
	}

	cfg.MinIOAccessKey = os.Getenv("MINIO_ACCESS_KEY")
	cfg.MinIOSecretKey = os.Getenv("MINIO_SECRET_KEY")
	cfg.MinIOBucket = os.Getenv("MINIO_BUCKET")
	if cfg.MinIOBucket == "" {
		cfg.MinIOBucket = "img"
	}
	// MINIO_USE_SSL: any non-empty and not equal to "false" or "0" means true
	if v := os.Getenv("MINIO_USE_SSL"); v != "" && v != "false" && v != "0" {
		cfg.MinIOUseSSL = true
	}

	log.Info("config parsed")

	return cfg, nil
}
