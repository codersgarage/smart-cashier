package data

import (
	"github.com/codersgarage/smart-cashier/models"
	"github.com/jinzhu/gorm"
)

type DiaryRepositoryImpl struct {
}

var diaryRepository DiaryRepository

func NewDiaryRepository() DiaryRepository {
	if diaryRepository == nil {
		diaryRepository = &DiaryRepositoryImpl{}
	}

	return diaryRepository
}

func (dr *DiaryRepositoryImpl) CreateDiary(db *gorm.DB, d *models.Diary) error {
	return db.Table(d.TableName()).Create(d).Error
}

func (dr *DiaryRepositoryImpl) ListDiaries(db *gorm.DB, userID string, from, limit int) ([]models.Diary, error) {
	var diaries []models.Diary
	d := models.Diary{}
	if err := db.Table(d.TableName()).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(from).
		Find(&diaries).Error; err != nil {
		return nil, err
	}
	return diaries, nil
}

func (dr *DiaryRepositoryImpl) SearchDiaries(db *gorm.DB, query, userID string, from, limit int) ([]models.Diary, error) {
	var diaries []models.Diary
	d := models.Diary{}
	if err := db.Table(d.TableName()).
		Where("user_id = ? AND name LIKE ?", userID, "%"+query+"%").
		Order("created_at DESC").
		Limit(limit).
		Offset(from).
		Find(&diaries).Error; err != nil {
		return nil, err
	}
	return diaries, nil
}

func (dr *DiaryRepositoryImpl) DeleteDiary(db *gorm.DB, userID, diaryID string) error {
	d := models.Diary{}
	if err := db.Table(d.TableName()).
		Where("user_id = ? AND id = ?", userID, diaryID).
		Delete(&d).Error; err != nil {
		return err
	}
	return nil
}

func (dr *DiaryRepositoryImpl) GetDiary(db *gorm.DB, userID, diaryID string) (*models.Diary, error) {
	d := models.Diary{}
	if err := db.Table(d.TableName()).
		Where("user_id = ? AND id = ?", userID, diaryID).
		Find(&d).Error; err != nil {
		return nil, err
	}
	return &d, nil
}

func (dr *DiaryRepositoryImpl) UpdateDiary(db *gorm.DB, d *models.Diary) error {
	if err := db.Table(d.TableName()).
		Where("user_id = ? AND id = ?", d.UserID, d.ID).
		Select("name").
		Update(map[string]interface{}{
			"name": d.Name,
		}).Error; err != nil {
		return err
	}
	return nil
}

//func (dr *DiaryRepositoryImpl) CreateEntry(db *gorm.DB, e *models.Entry) error {
//	return db.Table(e.TableName()).Create(e).Error
//}
//
//func (dr *DiaryRepositoryImpl) ListEntries(db *gorm.DB, userID, diaryID string, from, limit int) ([]models.Entry, error) {
//	var entries []models.Entry
//	e := models.Entry{}
//	if err := db.Table(e.TableName()).
//		Where("diary_id = ?", userID, diaryID).
//		Limit(limit).
//		Offset(from).
//		Find(&entries).Error; err != nil {
//		return nil, err
//	}
//	return entries, nil
//}
//
//func (dr *DiaryRepositoryImpl) SearchEntries(db *gorm.DB, query, userID, diaryID string, from, limit int) ([]models.Entry, error) {
//	var entries []models.Entry
//	e := models.Entry{}
//	if err := db.Table(e.TableName()).
//		Where("user_id = ? AND diary_id = ? AND name LIKE = ?", userID, diaryID, "%"+query+"%").
//		Limit(limit).
//		Offset(from).
//		Find(&entries).Error; err != nil {
//		return nil, err
//	}
//	return entries, nil
//}
//
//func (dr *DiaryRepositoryImpl) DeleteEntry(db *gorm.DB, userID, diaryID, entryID string) error {
//	e := models.Entry{}
//	if err := db.Table(e.TableName()).
//		Where("user_id = ? AND diary_id = ? AND id = ?", userID, diaryID, entryID).
//		Delete(&e).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func (dr *DiaryRepositoryImpl) GetEntry(db *gorm.DB, userID, diaryID, entryID string) (*models.Entry, error) {
//	e := models.Entry{}
//	if err := db.Table(e.TableName()).
//		Where("user_id = ? AND diary_id = ? AND id = ?", userID, diaryID, entryID).
//		Find(&e).Error; err != nil {
//		return nil, err
//	}
//	return &e, nil
//}
//
//func (dr *DiaryRepositoryImpl) UpdateEntry(db *gorm.DB, e *models.Entry) error {
//	e := models.Entry{}
//	if err := db.Table(e.TableName()).
//		Where("user_id = ? AND diary_id = ? AND id = ?", e.DiaryID, diaryID, entryID).
//		Find(&e).Error; err != nil {
//		return nil, err
//	}
//	return &e, nil
//}
//
//func (dr *DiaryRepositoryImpl) CreateEntryCategory(db *gorm.DB, ec *models.Category) error {
//
//}
//
//func (dr *DiaryRepositoryImpl) ListEntryCategories(db *gorm.DB, userID, diaryID string, from, limit int) ([]models.Category, error) {
//
//}
//
//func (dr *DiaryRepositoryImpl) SearchEntryCategories(db *gorm.DB, query, userID, diaryID string, from, limit int) ([]models.Category, error) {
//
//}
//
//func (dr *DiaryRepositoryImpl) DeleteEntryCategory(db *gorm.DB, userID, diaryID, entryCategoryID string) error {
//
//}
//
//func (dr *DiaryRepositoryImpl) GetEntryCategory(db *gorm.DB, userID, diaryID, entryCategoryID string) (*models.Category, error) {
//
//}
//
//func (dr *DiaryRepositoryImpl) UpdateEntryCategory(db *gorm.DB, ec *models.Category) error {
//
//}
