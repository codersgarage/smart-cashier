package data

import (
	"fmt"
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
		Select("name, type").
		Update(map[string]interface{}{
			"name": d.Name,
			"type": d.Type,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (dr *DiaryRepositoryImpl) CreateEntry(db *gorm.DB, e *models.Entry) error {
	return db.Table(e.TableName()).Create(e).Error
}

func (dr *DiaryRepositoryImpl) ListEntries(db *gorm.DB, userID, diaryID string, from, limit int) ([]models.EntryDetails, error) {
	var entries []models.EntryDetails
	e := models.Entry{}
	d := models.Diary{}
	c := models.Category{}
	if err := db.Table(fmt.Sprintf("%s AS e", e.TableName())).
		Select("e.id AS id, e.amount AS amount, e.note AS note, e.created_at AS created_at, e.updated_at AS updated_at,"+
			" d.id AS diary_id, d.name AS diary_name, d.user_id AS user_id, c.id AS category_id, c.name AS category_name").
		Joins(fmt.Sprintf("LEFT JOIN %s AS d ON e.diary_id = d.id", d.TableName())).
		Joins(fmt.Sprintf("LEFT JOIN %s AS c ON e.category_id = c.id", c.TableName())).
		Where("d.user_id = ? AND e.diary_id = ?", userID, diaryID).
		Order("created_at DESC").
		Limit(limit).
		Offset(from).
		Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (dr *DiaryRepositoryImpl) DeleteEntry(db *gorm.DB, diaryID, entryID string) error {
	e := models.Entry{}
	if err := db.Table(e.TableName()).
		Where("diary_id = ? AND id = ?", diaryID, entryID).
		Delete(&e).Error; err != nil {
		return err
	}
	return nil
}

func (dr *DiaryRepositoryImpl) DeleteAllEntry(db *gorm.DB, diaryID string) error {
	e := models.Entry{}
	if err := db.Table(e.TableName()).
		Where("diary_id = ?", diaryID).
		Delete(&e).Error; err != nil {
		return err
	}
	return nil
}

func (dr *DiaryRepositoryImpl) GetEntry(db *gorm.DB, userID, diaryID, entryID string) (*models.EntryDetails, error) {
	entry := models.EntryDetails{}
	e := models.Entry{}
	d := models.Diary{}
	c := models.Category{}
	if err := db.Table(fmt.Sprintf("%s AS e", e.TableName())).
		Select("e.id AS id, e.amount AS amount, e.note AS note, e.created_at AS created_at, e.updated_at AS updated_at,"+
			" d.id AS diary_id, d.name AS diary_name, d.user_id AS user_id, c.id AS category_id, c.name AS category_name").
		Joins(fmt.Sprintf("LEFT JOIN %s AS d ON e.diary_id = d.id", d.TableName())).
		Joins(fmt.Sprintf("LEFT JOIN %s AS c ON e.category_id = c.id", c.TableName())).
		Where("d.user_id = ? AND e.diary_id = ? AND e.id = ?", userID, diaryID, entryID).
		Find(&entry).Error; err != nil {
		return nil, err
	}
	return &entry, nil
}

func (dr *DiaryRepositoryImpl) UpdateEntry(db *gorm.DB, e *models.Entry) error {
	if err := db.Table(e.TableName()).
		Where("diary_id = ? AND id = ?", e.DiaryID, e.ID).
		Find(&e).Error; err != nil {
		return err
	}
	return nil
}

func (dr *DiaryRepositoryImpl) CreateCategory(db *gorm.DB, c *models.Category) error {
	return db.Table(c.TableName()).Create(c).Error
}

func (dr *DiaryRepositoryImpl) ListCategories(db *gorm.DB, userID, diaryID string, from, limit int) ([]models.CategoryDetails, error) {
	var categories []models.CategoryDetails
	c := models.Category{}
	d := models.Diary{}
	e := models.Entry{}
	if err := db.Table(fmt.Sprintf("%s AS c", c.TableName())).
		Select("c.id AS id, c.name AS name, SUM(e.amount) AS amount").
		Joins(fmt.Sprintf("JOIN %s AS d ON c.diary_id = d.id", d.TableName())).
		Joins(fmt.Sprintf("LEFT JOIN %s AS e ON c.id = e.category_id", e.TableName())).
		Group("c.id, c.name, c.created_at").
		Where("d.user_id = ? AND c.diary_id = ?", userID, diaryID).
		Order("c.created_at DESC").
		Limit(limit).
		Offset(from).
		Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (dr *DiaryRepositoryImpl) SearchCategories(db *gorm.DB, query, userID, diaryID string, from, limit int) ([]models.Category, error) {
	var categories []models.Category
	c := models.Category{}
	if err := db.Table(c.TableName()).
		Where("user_id = ? AND id = ? AND name LIKE ?", userID, diaryID, "%"+query+"%").
		Order("created_at DESC").
		Limit(limit).
		Offset(from).
		Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (dr *DiaryRepositoryImpl) DeleteCategory(db *gorm.DB, diaryID, categoryID string) error {
	c := models.Category{}
	if err := db.Table(c.TableName()).
		Where("diary_id = ? AND id = ?", diaryID, categoryID).
		Delete(&c).Error; err != nil {
		return err
	}
	return nil
}

func (dr *DiaryRepositoryImpl) DeleteAllCategory(db *gorm.DB, diaryID string) error {
	c := models.Category{}
	if err := db.Table(c.TableName()).
		Where("diary_id = ?", diaryID).
		Delete(&c).Error; err != nil {
		return err
	}
	return nil
}

func (dr *DiaryRepositoryImpl) GetCategory(db *gorm.DB, userID, diaryID, categoryID string) (*models.Category, error) {
	c := models.Category{}
	if err := db.Table(c.TableName()).
		Where("user_id = ? AND id = ?", userID, diaryID).
		Find(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (dr *DiaryRepositoryImpl) UpdateCategory(db *gorm.DB, c *models.Category) error {
	if err := db.Table(c.TableName()).
		Select("name").
		Update(map[string]interface{}{
			"name": c.Name,
		}).
		Where("id = ?", c.ID).
		Find(&c).Error; err != nil {
		return err
	}
	return nil
}
