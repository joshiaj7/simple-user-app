package model

// User is table model for user table
type User struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"user_id"`
	UserName   string `json:"user_name"`
	UUID       string `json:"uuid"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Password   string `json:"password"`
	IsLoggedIn bool   `json:"is_logged_in"`
}
