package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetDBConnection() (*sql.DB, error) {
	// Get the SQL Server connection details from environment variables
	server := os.Getenv("SQL_SERVER")
	port := os.Getenv("SQL_PORT")
	user := os.Getenv("SQL_USER")
	password := os.Getenv("SQL_PASSWORD")
	database := os.Getenv("SQL_DATABASE")

	// Create the connection string
	connString := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s",
		server, port, user, password, database)

	// Open a connection to the SQL Server database
	db, err := sql.Open("mssql", connString)
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
