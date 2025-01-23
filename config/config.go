package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Server struct {
	Port              string        `mapstructure:"PORT"`
	ReadTimeout       time.Duration `mapstructure:"READ_TIMEOUT"`
	WriteTimeout      time.Duration `mapstructure:"WRITE_TIMEOUT"`
	IdleTimeout       time.Duration `mapstructure:"IDLE_TIMEOUT"`
	ShutdownGraceTime time.Duration `mapstructure:"SHUTDOWN_GRACE_TIME"`
}

type Metadata struct {
	FilePath         string        `mapstructure:"METADATA_FILE_PATH"`
	AutoSaveInterval time.Duration `mapstructure:"AUTO_SAVE_INTERVAL"`
}

type Logger struct {
	Level string `mapstructure:"LOG_LEVEL"`
}

type Config struct {
	Server   Server
	Metadata Metadata
	Logger   Logger
}

func LoadConfig() (*Config, error) {
	setDefaults()

	if err := LoadEnvFile(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	viper.AutomaticEnv()

	if err := validateRequiredFields(); err != nil {
		return nil, err
	}

	var cfg Config

	cfg.Server.Port = viper.GetString("PORT")
	cfg.Server.ReadTimeout = viper.GetDuration("READ_TIMEOUT") * time.Second
	cfg.Server.WriteTimeout = viper.GetDuration("WRITE_TIMEOUT") * time.Second
	cfg.Server.IdleTimeout = viper.GetDuration("IDLE_TIMEOUT") * time.Second
	cfg.Server.ShutdownGraceTime = viper.GetDuration("SHUTDOWN_GRACE_TIME") * time.Second
	cfg.Metadata.FilePath = viper.GetString("METADATA_FILE_PATH")
	cfg.Metadata.AutoSaveInterval = viper.GetDuration("AUTO_SAVE_INTERVAL") * time.Second
	cfg.Logger.Level = viper.GetString("LOG_LEVEL")

	return &cfg, nil
}

func setDefaults() {
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("SWAGGER_ROUTE", "/swagger/*any")
	viper.SetDefault("READ_TIMEOUT", 5)        // seconds
	viper.SetDefault("WRITE_TIMEOUT", 10)      // seconds
	viper.SetDefault("IDLE_TIMEOUT", 60)       // seconds
	viper.SetDefault("SHUTDOWN_GRACE_TIME", 5) // seconds
	viper.SetDefault("METADATA_FILE_PATH", "./metadata.json")
	viper.SetDefault("AUTO_SAVE_INTERVAL", 300) // seconds
	viper.SetDefault("LOG_LEVEL", "info")
}

func validateRequiredFields() error {
	required := map[string]string{
		"METADATA_FILE_PATH": "path to metadata storage",
		"PORT":               "server port",
	}

	for field, desc := range required {
		if viper.GetString(field) == "" {
			return fmt.Errorf("required configuration missing: %s (%s)", field, desc)
		}
	}
	return nil
}

func LoadEnvFile() error {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(*os.PathError); ok {
			return nil
		}
		return fmt.Errorf("error reading .env file: %w", err)
	}
	return nil
}