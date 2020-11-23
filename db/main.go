package db

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/kamaal111/color-selector-api/models"
)

// Connect to database
func Connect() *pg.DB {
	options := &pg.Options{
		User:     "postgres",
		Password: "pass",
		Addr:     "127.0.0.1:5432",
	}

	pgDB := pg.Connect(options)
	if pgDB == nil {
		log.Println("Failed to connect to database.")
		os.Exit(100)
	}
	log.Println("Connection to database successful.")

	createSchemaErr := createSchema(pgDB)
	if createSchemaErr != nil {
		log.Printf("Error while creating schema, Reason: %v\n", createSchemaErr)
		os.Exit(100)
	}
	log.Println("Successfully created database schemas")

	return pgDB
}

func createSchema(pgDB *pg.DB) error {
	createUserErr := models.CreateUsersTable(pgDB)
	if createUserErr != nil {
		return createUserErr
	}
	createSavedColorErr := models.CreateSavedColorsTable(pgDB)
	if createSavedColorErr != nil {
		return createSavedColorErr
	}
	return nil
}
