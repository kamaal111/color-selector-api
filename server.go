package main

import (
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/v10"

	"github.com/kamaal111/color-selector-api/db"
	"github.com/kamaal111/color-selector-api/models"
)

func main() {
	start := time.Now()

	pgDB := db.Connect()

	// if saveErr := saveDummySavedColor(pgDB); saveErr != nil {
	// 	log.Printf("Error while saving user, Reason: %v\n", saveErr)
	// 	os.Exit(100)
	// }

	// saveErr := saveDummyUser(pgDB)
	// if saveErr != nil {
	// 	log.Printf("Error while saving user, Reason: %v\n", saveErr)
	// 	os.Exit(100)
	// }

	user, getUserErr := models.GetUserByEmail(pgDB, "dummy@dummmy.com")
	if getUserErr != nil {
		log.Printf("Error while saving user, Reason: %v\n", getUserErr)
		os.Exit(100)
	}
	log.Printf("%d", user.ID)

	closeErr := pgDB.Close()
	if closeErr != nil {
		log.Printf("Error while closing the connection, Reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Println("Connection to database closed successful.")

	elapsed := time.Since(start)
	log.Printf("Operation took %s", elapsed)
}

func saveDummySavedColor(pgDB *pg.DB, user *models.User) error {
	newSavedColor := models.SavedColor{
		Name:      "Some Name",
		Hex:       "#111111",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		User:      user,
	}
	saveErr := newSavedColor.Save(pgDB)
	return saveErr
}

func saveDummyUser(pgDB *pg.DB) error {
	newUser := models.User{
		Email:     "dummy2@dummmy.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	saveErr := newUser.Save(pgDB)
	return saveErr
}
