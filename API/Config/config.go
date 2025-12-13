package Config

import (
	//"fmt"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	WebAPP_Port                 string
	Chalas_Forum_Host           string
	Chalas_Forum_Port           string
	Chalas_Forum_Name           string
	Chalas_Forum_Admin_Username string
	Chalas_Forum_Admin_Password string
}

func ConfigLoad() (*Config, error) {
	err := godotenv.Load("pkg/.env")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return &Config{
		WebAPP_Port:                 os.Getenv("FORUM_PORT"),
		Chalas_Forum_Host:           os.Getenv("FORUM_DB_HOST"),
		Chalas_Forum_Port:           os.Getenv("FORUM_DB_PORT"),
		Chalas_Forum_Name:           os.Getenv("FORUM_DB_NAME"),
		Chalas_Forum_Admin_Username: os.Getenv("FORUM_ADMIN_USER"),
		Chalas_Forum_Admin_Password: os.Getenv("FORUM_ADMIN_PASSWORDS"),
	}, nil
}
