package model

// Token is table model for token table
type Token struct {
	UserID     uint `gorm:"primaryKey" json:"user_id"`
	IsLoggedIn bool `json:"is_logged_in"`
}
