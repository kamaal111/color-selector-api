package main

import (
	"log"
	"os"

	"github.com/kamaal111/color-selector-api/db"
	"github.com/kamaal111/color-selector-api/router"
)

func main() {
	pgDB := db.Connect()

	router.HandleRequests(pgDB)

	closeErr := pgDB.Close()
	if closeErr != nil {
		log.Printf("Error while closing the connection, Reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Println("Connection to database closed successful.")

}
