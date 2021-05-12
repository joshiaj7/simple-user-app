package controller

import (
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
		return
	}

	// decode request body
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		view.HTTPResponse(w, 500, err.Error(), nil)
		return
	}

	// get user
	hashed_password := util.Encrypt([]byte(user.Password))
	err = config.DB.Where("user_name = ? and password = ?", user.UserName, hashed_password).First(&user).Error
	if err != nil {
		view.HTTPResponse(w, 401, err.Error(), nil)
		return
	}

	// change user isloggedin to true
	user.IsLoggedIn = true
	config.DB.Save(&user)

	view.HTTPResponse(w, 200, "Successfully logged in", nil)
	return
}

// logout to logging out from app
func logout(w http.ResponseWriter, r *http.Request) {
	// method validation
	if r.Method != http.MethodPost {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
		return
	}

	// get user data from auth header
	status, user := util.CheckIfLoggedIn(w, r)
	if status == false {
		return
	}

	// change user IsLoggedIn to false
	user.IsLoggedIn = false
	config.DB.Save(&user)

	view.HTTPResponse(w, 200, "Successfully logged out", nil)
}

// createUser to insert user to db
func createUser(w http.ResponseWriter, r *http.Request) {
	// method validation
	if r.Method != http.MethodPost {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
		return
	}

	// parse request body
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		view.HTTPResponse(w, 500, err.Error(), nil)
		return
	}

	var existing_users []model.User
	config.DB.Where("email = ? OR user_name = ?", user.Email, user.UserName).Find(&existing_users)
	if len(existing_users) > 0 {
		view.HTTPResponse(w, 403, "User already exist", nil)
		return
	}

	// generate uuid & create user
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		view.HTTPResponse(w, 500, err.Error(), nil)
		return
	}
	user.UUID = string(uuid[:len(uuid)-2])
	user.Password = util.Encrypt([]byte(user.Password))
	user.IsLoggedIn = true
	config.DB.Create(&user)

	view.HTTPResponse(w, 200, "Success", user)
}

// updateUser to update user from db
func updateUser(w http.ResponseWriter, r *http.Request) {
	// method validation
	if r.Method != http.MethodPut {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
		return
	}

	// get db user data from auth token
	status, db_user := util.CheckIfLoggedIn(w, r)
	if status == false {
		return
	}

	// parse request body
	decoder := json.NewDecoder(r.Body)
	var inc_user model.User
	err := decoder.Decode(&inc_user)
	if err != nil {
		view.HTTPResponse(w, 500, err.Error(), nil)
		return
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
		return
	}

	// get user data from auth header
	status, _ := util.CheckIfLoggedIn(w, r)
	if status == false {
		return
	}

	// parse request body
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		view.HTTPResponse(w, 500, err.Error(), nil)
	}

	// delete user
	config.DB.Delete(&user, user.ID)

	view.HTTPResponse(w, 200, "User has been deleted", nil)
}

// getUser to select all user from db
func getUser(w http.ResponseWriter, r *http.Request) {
	// check for method
	if r.Method != http.MethodGet {
		view.HTTPResponse(w, 405, "Method is not allowed", nil)
		return
	}

	// get user data from auth header
	status, _ := util.CheckIfLoggedIn(w, r)
	if status == false {
		return
	}

	// fetch all user data
	var users []model.User
	config.DB.Find(&users)

	view.HTTPResponse(w, 200, "Success", users)
}
