package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/serediukit/dating-wars/model"
	"github.com/serediukit/dating-wars/util"
)

var users []model.User
var searchableUsers []model.SearchableUser

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")

	router.HandleFunc("/pickup", getSearchableUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	newUserData := model.UserData{}
	err := json.NewDecoder(r.Body).Decode(&newUserData)
	if err != nil {
		fmt.Println(err)
	}

	users = append(users, newUserData.User)
	searchableUsers = append(searchableUsers, newUserData.SearchableUser)

	err = json.NewEncoder(w).Encode(newUserData)
	if err != nil {
		fmt.Println(err)
	}

	_, err = w.Write(
		[]byte(
			fmt.Sprintf(
				"Country: %s\nGender: %s\nSearch Gender: %s\n",
				util.CountryNames[newUserData.SearchableUser.CountryID],
				util.Genders[newUserData.SearchableUser.GenderID],
				util.Genders[newUserData.SearchableUser.SearchGenderID],
			),
		),
	)
	if err != nil {
		fmt.Println(err)
	}
}

func getSearchableUsers(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(searchableUsers)
	if err != nil {
		fmt.Println(err)
	}
}
