package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
}

var users []User

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	newUser := User{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
	}

	users = append(users, newUser)
	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		fmt.Println(err)
	}
}
