package internalsql

import (
	"database/sql"
	"fmt"
	"go-tweets/internal/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(cfg *config.Config) (*sql.DB, error) {

	databaseSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, "Asia%2FKolkata")

	db, err := sql.Open("mysql", databaseSourceName)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database")
	}

	log.Println("database connected")
	return db, nil
}
