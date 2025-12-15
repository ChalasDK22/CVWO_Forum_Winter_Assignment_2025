package internalsql

import (
	"database/sql"
	"fmt"
	"log"

	"chalas.com/forum_project/API/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectAPI_MYSQL(chalasconfig *config.Config) (*sql.DB, error) {
	sourcename := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		chalasconfig.Chalas_Forum_Admin_Username,
		chalasconfig.Chalas_Forum_Admin_Password,
		chalasconfig.Chalas_Forum_Host,
		chalasconfig.WebAPP_Port,
		chalasconfig.Chalas_Forum_Name,
	)
	chalas_mydb, err := sql.Open("mysql", sourcename)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	log.Println("Connected")
	return chalas_mydb, nil
}
