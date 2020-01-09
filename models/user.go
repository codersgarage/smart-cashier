package models

import (
	"time"
)

const (
	UserRegistered UserStatus = "registered"
	UserActive     UserStatus = "active"
	UserBanned     UserStatus = "banned"
)

type UserStatus string

type User struct {
	ID                string     `json:"id" gorm:"column:id;primary_key"`
	Name              string     `json:"name" gorm:"column:name;not null"`
	Email             string     `json:"email" gorm:"column:email;unique;not null"`
	Password          string     `json:"-" gorm:"column:password;not null"`
	VerificationToken *string    `json:"-" gorm:"column:verification_token;unique"`
	Status            UserStatus `json:"status" gorm:"column:status;index;not null"`
	CreatedAt         time.Time  `json:"created_at" gorm:"column:created_at;index"`
	UpdatedAt         time.Time  `json:"updated_at" gorm:"column:updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
