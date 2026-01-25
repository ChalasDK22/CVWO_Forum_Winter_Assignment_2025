package config

import (
	//"fmt"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WebAPP_Port string
	//Chalas_Forum_Host           string
	//Chalas_Forum_Port           string
	//Chalas_Forum_Name           string
	//Chalas_Forum_Admin_Username string
	//Chalas_Forum_Admin_Password string
	Chalas_JWT    string
	Chalas_DB_Url string
}

func ConfigLoad() (*Config, error) {

	// Only load .env in LOCAL development
	if os.Getenv("ENV") != "production" {
		_ = godotenv.Load(".env") // ignore error on purpose
	}

	cfg := &Config{
		WebAPP_Port:   os.Getenv("FORUM_PORT"), // Render provides PORT
		Chalas_DB_Url: os.Getenv("DATABASE_URL"),
		Chalas_JWT:    os.Getenv("FORUM_JWT"),
	}

	// Fallback for local dev
	if cfg.WebAPP_Port == "" {
		cfg.WebAPP_Port = os.Getenv("FORUM_PORT")
	}
	if cfg.WebAPP_Port == "" {
		cfg.WebAPP_Port = "8080"
	}

	if cfg.Chalas_DB_Url == "" {
		return nil, fmt.Errorf("DATABASE_URL is missing")
	}

	return cfg, nil
}
