package data

import (
	"github.com/codersgarage/smart-cashier/models"
	"github.com/codersgarage/smart-cashier/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type UserRepositoryImpl struct {
}

var userRepository UserRepository

func NewUserRepository() UserRepository {
	if userRepository == nil {
		userRepository = &UserRepositoryImpl{}
	}
	return userRepository
}

func (uu *UserRepositoryImpl) Register(db *gorm.DB, u *models.User) error {
	if err := db.Model(u).Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (uu *UserRepositoryImpl) Login(db *gorm.DB, email, password string) (*models.Session, error) {
	u := models.User{}

	if err := db.Table(u.TableName()).Where("email = ? AND status = ?", email, models.UserActive).Find(&u).Error; err != nil {
		return nil, err
	}

	if err := utils.CheckPassword(u.Password, password); err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	s := models.Session{
		ID:           utils.NewUUID(),
		UserID:       u.ID,
		AccessToken:  utils.NewToken(),
		RefreshToken: utils.NewToken(),
		CreatedAt:    time.Now().UTC(),
		ExpireOn:     time.Now().Add(time.Hour * 48).Unix(),
	}

	if err := db.Model(&s).Create(&s).Error; err != nil {
		return nil, err
	}

	return &s, nil
}

func (uu *UserRepositoryImpl) Logout(db *gorm.DB, token string) error {
	s := models.Session{}

	if err := db.Model(&s).Where("access_token = ?", token).Delete(&s).Error; err != nil {
		return err
	}
	return nil
}

func (uu *UserRepositoryImpl) RefreshToken(db *gorm.DB, token string) (*models.Session, error) {
	os := models.Session{}

	if err := db.Model(&os).Where("refresh_token = ?", token).First(&os).Error; err != nil {
		return nil, err
	}

	s := models.Session{
		ID:           utils.NewUUID(),
		UserID:       os.UserID,
		AccessToken:  utils.NewToken(),
		RefreshToken: utils.NewToken(),
		CreatedAt:    time.Now().UTC(),
		ExpireOn:     time.Now().Add(time.Hour * 48).Unix(),
	}

	if err := db.Model(&s).Create(&s).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&os).Where("refresh_token = ?", token).Delete(&os).Error; err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}

	return &s, nil
}

func (uu *UserRepositoryImpl) Update(db *gorm.DB, u *models.User) error {
	if err := db.Table(u.TableName()).
		Where("id = ?", u.ID).
		Select("name, profile_picture, phone, password, verification_token, is_email_verified, status, permission_id, updated_at").
		Updates(map[string]interface{}{
			"name":               u.Name,
			"profile_picture":    u.ProfilePicture,
			"password":           u.Password,
			"verification_token": u.VerificationToken,
			"status":             u.Status,
			"updated_at":         u.UpdatedAt,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (uu *UserRepositoryImpl) Get(db *gorm.DB, userID string) (*models.User, error) {
	u := models.User{}

	if err := db.Model(&u).Where("id = ?", userID).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (uu *UserRepositoryImpl) GetSession(db *gorm.DB, token string) (*models.Session, error) {
	s := models.Session{}

	if err := db.Table(s.TableName()).Where("access_token = ?", token).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}
