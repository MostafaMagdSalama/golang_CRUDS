package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	fmt.Println("hello")
	db, err := sql.Open("postgress", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("error in open DB connection")
	}

	defer db.Close()
	//   handle router
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", getUser(db)).Methods("GET")
	router.HandleFunc("/user", createUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", updateuser(db)).Methods("PUT")

	// start server
	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")
		next.ServeHTTP(w,r)

	})
}
