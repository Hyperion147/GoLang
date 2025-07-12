package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id     string    `json:"user_id"`
	Name   string `json:"username"`
	Device string `json:"user_device"`
	Admin  *Admin `json:"admin_user"`
}

type Admin struct {
	AdminId   string    `json:"admin_id"`
	AdminName string `json:"admin_name"`
}

var users []User

func (u *User) CheckIsEmpty() bool {
	return u.Name == "" && u.Id == ""
}

func main() {
	topic := "This program is to learn router."
	fmt.Println(topic)

	r := mux.NewRouter()
	r.HandleFunc("/", metaverseHome).Methods("GET")
	r.HandleFunc("/users", getAllUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))

}

func metaverseHome (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Test api for metaverseTwoDee</h1>"))
}

func getAllUsers (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all users")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a user")
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)
	fmt.Println(params)
	fmt.Printf("%T",params)

	for _, user := range users{
		if user.Id == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}	
	}

	json.NewEncoder(w).Encode("No user found with given id.")

}