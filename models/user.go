package models

import (
	"log"
	"time"

	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Fullname  string    `gorm:"type:varchar(255);not null" json:"fullname"`
	Username  string    `gorm:"type:varchar(255);not null;unique" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null"`
	HasAdmin  bool      `gorm:"type:bool;default:false" json:"has_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// hook
func (user *User) BeforeCreate(db *gorm.DB) error {
	user.Password = hashPassword(user.Password)
	return nil
}

// func
func hashPassword(password string) string {
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		log.Fatal(err)
	}
	return string(encoded)
}
