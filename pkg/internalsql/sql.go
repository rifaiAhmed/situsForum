package internalsql

import (
	"database/sql"
	"log"
)

func Connect(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("error connecting to databse")
		return nil, err
	}
	return db, nil
}
