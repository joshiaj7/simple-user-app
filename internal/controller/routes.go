package controller

import (
	"github.com/joshiaj7/simple-user-app/internal/config"
	"github.com/joshiaj7/simple-user-app/internal/model"
	"github.com/joshiaj7/simple-user-app/internal/util"
	"github.com/joshiaj7/simple-user-app/internal/view"

	"encoding/json"
	"fmt"
	"net/http"
)

var key = "User-Simple-App"

// SetupRoute to handle routing
func SetupRoute() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/create", createUser)
	http.HandleFunc("/update", updateUser)
	http.HandleFunc("/delete", deleteUser)
	http.HandleFunc("/get", getUser)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is login")

	// check for method
	if r.Method != http.MethodPost {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
	}

	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user.Email, user.Password)
	fmt.Println("hashed password : ", util.Encrypt([]byte(user.Password), key))

	view.HTTPResponse(w, 200, "Successfully logged in", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is logout")

	// check for method
	if r.Method != http.MethodPost {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
	}

	view.HTTPResponse(w, 200, "Successfully logged out", nil)
}

// createUser to insert user to db
func createUser(w http.ResponseWriter, r *http.Request) {
	// check for method
	if r.Method != http.MethodPost {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
	}

	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	// pass pointer of user to Create
	config.DB.Create(&user)

	view.HTTPResponse(w, 200, "Success", user)
}

// updateUser to update user from db
func updateUser(w http.ResponseWriter, r *http.Request) {
	// check for method
	if r.Method != http.MethodPut {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
	}

	decoder := json.NewDecoder(r.Body)
	var inc_user, db_user model.User
	err := decoder.Decode(&inc_user)
	if err != nil {
		panic(err)
	}

	config.DB.First(&db_user, inc_user.ID)
	db_user.Email = inc_user.Email
	db_user.Address = inc_user.Address
	db_user.Password = inc_user.Password
	config.DB.Save(&db_user)

	view.HTTPResponse(w, 200, "User has been updated", db_user)
}

// deleteUser to delete user from db
func deleteUser(w http.ResponseWriter, r *http.Request) {
	// check for method
	if r.Method != http.MethodDelete {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
	}

	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	var token model.Token
	token.UserID = user.ID

	config.DB.Delete(&user, user.ID)
	config.DB.Delete(&token, token.UserID)

	view.HTTPResponse(w, 200, "User has been deleted", nil)
}

// getUser to select all user from db
func getUser(w http.ResponseWriter, r *http.Request) {
	// check for method
	if r.Method != http.MethodGet {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
	}

	var users []model.User
	config.DB.Find(&users)

	view.HTTPResponse(w, 200, "Success", users)
}
