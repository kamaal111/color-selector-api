package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/kamaal111/color-selector-api/models"
)

// HandleRequests is the main request handler
func HandleRequests(pgDB *pg.DB) {
	postSignUpUser := loggerMiddleware(connectToDatabase(pgDB, signUpUser))

	mux := http.NewServeMux()

	mux.Handle("/user", postSignUpUser)

	log.Println("Listening on :8080...")
	err := http.ListenAndServe("127.0.0.1:8080", mux)
	log.Fatal(err)
}

func signUpUser(w http.ResponseWriter, r *http.Request, pgDB *pg.DB) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var payload struct {
		Email string `json:"email"`
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		errorHandler(w, err.Error(), 400)
		return
	}

	if len(payload.Email) < 1 {
		errorHandler(w, "email has not been provided", 400)
		return
	}

	foundUser, getUserErr := models.GetUserByEmail(pgDB, payload.Email)
	if getUserErr == nil {
		errorHandler(w, fmt.Sprintf("user with the email %s allready exists", foundUser.Email), 409)
		return
	}

	user := models.User{
		Email:     payload.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	insertErr := user.Save(pgDB)
	if insertErr != nil {
		errorHandler(w, insertErr.Error(), 400)
		return
	}

	output, err := json.Marshal(payload)
	if err != nil {
		errorHandler(w, err.Error(), 400)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func connectToDatabase(pgDB *pg.DB, f func(w http.ResponseWriter, r *http.Request, pgDB *pg.DB)) http.Handler {
	funcToPass := func(w http.ResponseWriter, r *http.Request) {
		f(w, r, pgDB)
	}
	return http.HandlerFunc(funcToPass)
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		log.Printf("%s: %s in %s\n", r.Method, r.URL.Path, elapsed)
	})
}

// Error ...
type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func errorHandler(w http.ResponseWriter, error string, code int) {
	errorResponse := Error{
		Message: error,
		Status:  code,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(errorResponse.Status)
	json.NewEncoder(w).Encode(errorResponse)
}
