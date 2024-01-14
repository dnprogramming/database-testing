package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBConnection() (*sql.DB, error) {
	// Get the environment variables for the database connection
	host := os.Getenv("MARIADB_HOST")
	port := os.Getenv("MARIADB_PORT")
	user := os.Getenv("MARIADB_USER")
	password := os.Getenv("MARIADB_PASSWORD")
	database := os.Getenv("MARIADB_DATABASE")

	// Create the connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	// Open a connection to the database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
