package internalsql

import (
	"database/sql"
	"fmt"
	"log"

	"chalas.com/forum_project/API/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectAPI_MYSQL(chalasconfig *config.Config) (*sql.DB, error) {
	chalas_mydb, err := sql.Open("mysql", chalasconfig.Chalas_DB_Url)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	log.Println("Connected")
	return chalas_mydb, nil
}
