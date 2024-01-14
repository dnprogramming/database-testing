package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/ibmdb/go_ibm_db"
)

func GetInformixConnection() (*sql.DB, error) {
	// Get the environment variables
	host := os.Getenv("INFORMIX_HOST")
	port := os.Getenv("INFORMIX_PORT")
	user := os.Getenv("INFORMIX_USER")
	password := os.Getenv("INFORMIX_PASSWORD")
	database := os.Getenv("INFORMIX_DATABASE")

	// Create the connection string
	connStr := fmt.Sprintf("HOSTNAME=%s;PORT=%s;UID=%s;PWD=%s;DATABASE=%s;", host, port, user, password, database)

	// Open a connection to the Informix database
	db, err := sql.Open("go_ibm_db", connStr)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
