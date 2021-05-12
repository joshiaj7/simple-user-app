package model

// User is table model for user table
type User struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"user_id"`
	UserName   string `gorm:"unique" json:"user_name"`
	UUID       string `gorm:"unique" json:"uuid"`
	Email      string `gorm:"unique" json:"email"`
	Address    string `json:"address"`
	Password   string `json:"-"`
	IsLoggedIn bool   `json:"is_logged_in"`
}
