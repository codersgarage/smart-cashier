package data

import (
	"github.com/codersgarage/smart-cashier/models"
	"github.com/jinzhu/gorm"
)

type DiaryRepository interface {
	CreateDiary(db *gorm.DB, d *models.Diary) error
	ListDiaries(db *gorm.DB, userID string, from, limit int) ([]models.Diary, error)
	SearchDiaries(db *gorm.DB, query, userID string, from, limit int) ([]models.Diary, error)
	DeleteDiary(db *gorm.DB, userID, diaryID string) error
	GetDiary(db *gorm.DB, userID, diaryID string) (*models.Diary, error)
	UpdateDiary(db *gorm.DB, d *models.Diary) error

	CreateEntry(db *gorm.DB, e *models.Entry) error
	ListEntries(db *gorm.DB, userID, diaryID string, from, limit int) ([]models.EntryDetails, error)
	//SearchEntries(db *gorm.DB, query, userID, diaryID string, from, limit int) ([]models.Entry, error)
	//DeleteEntry(db *gorm.DB, userID, diaryID, entryID string) error
	//GetEntry(db *gorm.DB, userID, diaryID, entryID string) (*models.Entry, error)
	//UpdateEntry(db *gorm.DB, e *models.Entry) error
	//
	CreateCategory(db *gorm.DB, ec *models.Category) error
	ListCategories(db *gorm.DB, userID, diaryID string, from, limit int) ([]models.Category, error)
	SearchCategories(db *gorm.DB, query, userID, diaryID string, from, limit int) ([]models.Category, error)
	DeleteCategory(db *gorm.DB, userID, diaryID, entryCategoryID string) error
	GetCategory(db *gorm.DB, userID, diaryID, entryCategoryID string) (*models.Category, error)
	UpdateCategory(db *gorm.DB, ec *models.Category) error
}
