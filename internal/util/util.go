package util

import (
	"net/http"

	"github.com/joshiaj7/simple-user-app/internal/config"
	"github.com/joshiaj7/simple-user-app/internal/model"
	"github.com/joshiaj7/simple-user-app/internal/view"

	b64 "encoding/base64"
	"strings"
)

// Encrypt function used to encrypt password
func Encrypt(data []byte) string {
	return b64.StdEncoding.EncodeToString([]byte(data))
}

// Decrypt function used to decrypt password
func Decrypt(data string) string {
	sDec, err := b64.StdEncoding.DecodeString(data)
	if err != nil {
		return "Failed to decode string, error: " + err.Error()
	}

	return string(sDec)
}

func CheckIfLoggedIn(w http.ResponseWriter, r *http.Request) (bool, model.User) {
	auth := r.Header.Get("Authorization")

	if auth == "" {
		view.HTTPResponse(w, 401, "Unauthorized", nil)
	}

	uuid := strings.Split(auth, " ")[1]

	// get user by uuid (bearer)
	var user model.User
	err := config.DB.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		view.HTTPResponse(w, 401, "Unauthorized", nil)
		return false, user
	}

	if user.IsLoggedIn == false {
		view.HTTPResponse(w, 404, "User is not logged in", nil)
		return false, user
	}

	return true, user
}
