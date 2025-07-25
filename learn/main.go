package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Id     string `json:"user_id"`
	Name   string `json:"username"`
	Device string `json:"user_device"`
	Admin  *Admin `json:"admin_user"`
}

type Admin struct {
	AdminId   string `json:"admin_id"`
	AdminName string `json:"admin_name"`
}

var users []User

func (u *User) CheckIsEmpty() bool {
	return u.Name == ""
}

func main() {
	topic := "This program is to learn mux router."
	fmt.Println(topic)

	users = append(users, User{Id: "1001", Name: "Hyper", Device: "PC", Admin: &Admin{AdminId: "101", AdminName: "Bgs"}})
	users = append(users, User{Id: "1002", Name: "Hyperion", Device: "PC", Admin: &Admin{AdminId: "102", AdminName: "Dks"}})

	r := mux.NewRouter()

	r.HandleFunc("/", metaverseHome).Methods("GET")
	r.HandleFunc("/users", getAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
	r.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func metaverseHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Test api for metaverseTwoDee</h1>"))
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all users")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a user")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)
	fmt.Printf("%T", params)

	for _, user := range users {
		if user.Id == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	json.NewEncoder(w).Encode("No user found with given id.")

}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a user")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Enter some data.")
		return
	}

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.CheckIsEmpty() {
		json.NewEncoder(w).Encode("No user data")
		return
	}
	for _, u := range users {
		if u.Name == user.Name {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode("Username already taken")
			return
		}
	}
	for {
		user.Id = strconv.Itoa(rand.Intn(1000))
		exists := false
		for _, u := range users {
			if u.Id == user.Id {
				exists = true
				break
			}
		}
		if !exists {
			break
		}
	}
	users = append(users, user)
	json.NewEncoder(w).Encode(user)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a user")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for id, user := range users {
		if user.Id == params["id"] {
			users = append(users[:id], users[id+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.Id = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
		}
	}

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete User")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for id, user := range users {
		if user.Id == params["id"] {
			users = append(users[:id], users[id+1:]...)
			json.NewEncoder(w).Encode("User deleted")
			break
		}
	}

}
