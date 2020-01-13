package data

import (
	"github.com/codersgarage/smart-cashier/models"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Register(db *gorm.DB, u *models.User) error
	Login(db *gorm.DB, email, password string) (*models.Session, error)
	Logout(db *gorm.DB, token string) error
	RefreshToken(db *gorm.DB, token string) (*models.Session, error)
	Update(db *gorm.DB, u *models.User) error
	Get(db *gorm.DB, userID string) (*models.User, error)
	GetSession(db *gorm.DB, token string) (*models.Session, error)
}
