package controller

import (
	"fmt"

	"github.com/joshiaj7/simple-user-app/internal/config"
	"github.com/joshiaj7/simple-user-app/internal/model"
	"github.com/joshiaj7/simple-user-app/internal/util"
	"github.com/joshiaj7/simple-user-app/internal/view"

	"encoding/json"
	"net/http"
	"os/exec"
)

// SetupRoute to handle routing
func SetupRoute() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/create", createUser)
	http.HandleFunc("/update", updateUser)
	http.HandleFunc("/delete", deleteUser)
	http.HandleFunc("/get", getUser)
}

// login to logging in from app
func login(w http.ResponseWriter, r *http.Request) {
	// method validation
	if r.Method != http.MethodPost {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
	}

	// decode request body
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		view.HTTPResponse(w, 500, err.Error(), nil)
	}

	// check if user exists
	err = config.DB.First(&user, user.ID).Error
	if err != nil {
		view.HTTPResponse(w, 404, "Record not found", nil)
	}

	// change user isloggedin to true
	var userLogIn model.UserLogIn
	config.DB.First(&userLogIn, user.ID)
	userLogIn.IsLoggedIn = true
	config.DB.Save(&userLogIn)

	view.HTTPResponse(w, 200, "Successfully logged in", nil)
}

// logout to logging out from app
func logout(w http.ResponseWriter, r *http.Request) {
	// method validation
	if r.Method != http.MethodPost {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
	}

	fmt.Println("method clear")

	// get user data from auth header
	user := util.CheckIfLoggedIn(w, r)

	fmt.Println(user)

	// change user IsLoggedIn to false
	var userLogIn model.UserLogIn
	config.DB.First(&userLogIn, user.ID)
	userLogIn.IsLoggedIn = false
	config.DB.Save(&userLogIn)

	view.HTTPResponse(w, 200, "Successfully logged out", nil)
}

// createUser to insert user to db
func createUser(w http.ResponseWriter, r *http.Request) {
	// method validation
	if r.Method != http.MethodPost {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
	}

	// parse request body
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		view.HTTPResponse(w, 500, err.Error(), nil)
	}

	// generate uuid & create user
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		view.HTTPResponse(w, 500, err.Error(), nil)
	}
	user.UUID = string(uuid[:len(uuid)-2])
	user.Password = util.Encrypt([]byte(user.Password))
	config.DB.Create(&user)

	// create UserLogIn
	userLogIn := &model.UserLogIn{
		UserID:     user.ID,
		IsLoggedIn: true,
	}
	config.DB.Create(&userLogIn)

	view.HTTPResponse(w, 200, "Success", user)
}

// updateUser to update user from db
func updateUser(w http.ResponseWriter, r *http.Request) {
	// method validation
	if r.Method != http.MethodPut {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
	}

	// get db user data from auth token
	var inc_user, db_user model.User
	db_user = util.CheckIfLoggedIn(w, r)

	// parse request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&inc_user)
	if err != nil {
		view.HTTPResponse(w, 500, err.Error(), nil)
	}

	config.DB.First(&db_user, inc_user.ID)

	if db_user.Email != inc_user.Email && inc_user.Email != "" {
		db_user.Email = inc_user.Email
	}

	if db_user.UserName != inc_user.UserName && inc_user.UserName != "" {
		db_user.UserName = inc_user.UserName
	}

	if db_user.Address != inc_user.Address && inc_user.Address != "" {
		db_user.Address = inc_user.Address
	}

	if inc_user.Address != "" {
		enc_pass := util.Encrypt([]byte(inc_user.Password))
		if db_user.Password != enc_pass {
			db_user.Password = enc_pass
		}
	}
	config.DB.Save(&db_user)

	view.HTTPResponse(w, 200, "User has been updated", db_user)
}

// deleteUser to delete user from db
func deleteUser(w http.ResponseWriter, r *http.Request) {
	// check for method
	if r.Method != http.MethodDelete {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
	}

	// get user data from auth header
	_ = util.CheckIfLoggedIn(w, r)

	// parse request body
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		view.HTTPResponse(w, 500, err.Error(), nil)
	}

	fmt.Println("user: ", user)

	// delete user and user loggin
	var userLogIn model.UserLogIn
	userLogIn.UserID = user.ID

	config.DB.Delete(&user, user.ID)
	config.DB.Delete(&userLogIn, userLogIn.UserID)

	view.HTTPResponse(w, 200, "User has been deleted", nil)
}

// getUser to select all user from db
func getUser(w http.ResponseWriter, r *http.Request) {
	// check for method
	if r.Method != http.MethodGet {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
	}

	// get user data from auth header
	_ = util.CheckIfLoggedIn(w, r)

	// fetch all user data
	var users []model.User
	config.DB.Find(&users)

	view.HTTPResponse(w, 200, "Success", users)
}
