package models

import (
	"log"
	"time"

	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;not null"`
	Fullname  string `gorm:"type:varchar(255);not null"`
	Username  string `gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	HasAdmin  bool   `gorm:"type:bool;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// payload
type CreateUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
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
