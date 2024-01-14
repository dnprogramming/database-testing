package main

import (
	"context"

	db "github.com/dnprogramming/database-testing/go/helpers/technologies/database"
	"go.mongodb.org/mongo-driver/bson"
)

// Insert Config calls to the database and inserts the configuration
func InsertMSSQLConfig(config string) error {
	// Get the database connection
	database, err := db.GetDBConnection()
	if err != nil {
		return err
	}

	// Insert the configuration into the database
	_, err = database.Exec("INSERT INTO config (config) VALUES (?)", config)
	if err != nil {
		return err
	}

	return nil
}

func InsertInformixConfig(config string) error {
	// Get the database connection
	database, err := db.GetInformixConnection()
	if err != nil {
		return err
	}

	// Insert the configuration into the database
	_, err = database.Exec("INSERT INTO config (config) VALUES (?)", config)
	if err != nil {
		return err
	}

	return nil
}

func InsertMySQLConfig(config string) error {
	// Get the database connection
	database, err := db.NewMySQLDB()
	if err != nil {
		return err
	}

	// Insert the configuration into the database
	_, err = database.Exec("INSERT INTO config (config) VALUES (?)", config)
	if err != nil {
		return err
	}

	return nil
}

func InsertMariaDBConfig(config string) error {
	// Get the database connection
	database, err := db.NewDBConnection()
	if err != nil {
		return err
	}

	// Insert the configuration into the database
	_, err = database.Exec("INSERT INTO config (config) VALUES (?)", config)
	if err != nil {
		return err
	}

	return nil
}

func InsertPostgreSQLConfig(config string) error {
	// Get the database connection
	database, err := db.NewPostgresDB()
	if err != nil {
		return err
	}

	// Insert the configuration into the database
	_, err = database.Exec("INSERT INTO config (config) VALUES ($1)", config)
	if err != nil {
		return err
	}

	return nil
}

// Insert the configuration into the MongoDB Database as a document in a collection
func InsertMongoDBConfig(config string) error {
	// Get the MongoDB connection
	database, err := db.ConnectToMongoDB()
	if err != nil {
		return err
	}

	// Insert the configuration into the database
	_, err = database.Collection("config").InsertOne(context.Background(), bson.M{"config": config})
	if err != nil {
		return err
	}

	return nil
}

// Insert the configuration into the Apache Cassandra Database as a row in a table
func InsertCassandraConfig(config string) error {
	// Get the Cassandra connection
	database, err := db.NewCassandraDB()
	if err != nil {
		return err
	}

	// Insert the configuration into the database
	err = database.Query("INSERT INTO config (config) VALUES (?)", config).Exec()
	if err != nil {
		return err
	}

	return nil
}
