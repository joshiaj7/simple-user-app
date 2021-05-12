package model

// User is table model for user table
type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}
