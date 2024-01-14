package main

import (
	"context"

	db "github.com/dnprogramming/database-testing/go/helpers/technologies/database"
	"go.mongodb.org/mongo-driver/bson"
)

// GetConfig calls to the database and returns the configuration
func GetMSSQLConfig() (string, error) {
	// Get the database connection
	database, err := db.GetDBConnection()
	if err != nil {
		return "", err
	}

	// Query the database for the configuration
	var config string
	err = database.QueryRow("SELECT config FROM config").Scan(&config)
	if err != nil {
		return "", err
	}

	return config, nil
}

func GetInformixConfig() (string, error) {
	// Get the database connection
	database, err := db.GetInformixConnection()
	if err != nil {
		return "", err
	}

	// Query the database for the configuration
	var config string
	err = database.QueryRow("SELECT config FROM config").Scan(&config)
	if err != nil {
		return "", err
	}

	return config, nil
}

func GetMySQLConfig() (string, error) {
	// Get the database connection
	database, err := db.NewMySQLDB()
	if err != nil {
		return "", err
	}

	// Query the database for the configuration
	var config string
	err = database.QueryRow("SELECT config FROM config").Scan(&config)
	if err != nil {
		return "", err
	}

	return config, nil
}

func GetMariaDBConfig() (string, error) {
	// Get the database connection
	database, err := db.NewDBConnection()
	if err != nil {
		return "", err
	}

	// Query the database for the configuration
	var config string
	err = database.QueryRow("SELECT config FROM config").Scan(&config)
	if err != nil {
		return "", err
	}

	return config, nil
}

// GetConfig calls to the database and returns the configuration
func GetPostgresConfig() (string, error) {
	// Get the database connection
	database, err := db.NewPostgresDB()
	if err != nil {
		return "", err
	}

	// Query the database for the configuration
	var config string
	err = database.QueryRow("SELECT config FROM config").Scan(&config)
	if err != nil {
		return "", err
	}

	return config, nil
}

// GetConfig calls to the database and returns the configuration
func GetMongoDBConfig() (string, error) {
	// Get the database connection
	database, err := db.ConnectToMongoDB()
	if err != nil {
		return "", err
	}

	// Query the database for the configuration
	var config string
	err = database.Collection("config").FindOne(context.Background(), bson.M{}).Decode(&config)
	if err != nil {
		return "", err
	}

	return config, nil
}

// GetConfig calls to the database and returns the configuration
func GetCassandraConfig() (string, error) {
	// Get the database connection
	database, err := db.NewCassandraDB()
	if err != nil {
		return "", err
	}

	// Query the database for the configuration
	var config string
	err = database.Query("SELECT config FROM config").Scan(&config)
	if err != nil {
		return "", err
	}

	return config, nil
}
