package models

import (
	"fmt"
	"time"
)

type Entry struct {
	ID         string    `json:"id" gorm:"column:id;primary_key"`
	DiaryID    string    `json:"diary_id" gorm:"column:diary_id;index;not null"`
	CategoryID string    `json:"category_id" gorm:"column:category_id;index;not null"`
	Note       string    `json:"note" gorm:"column:note;not null"`
	Amount     int       `json:"amount" gorm:"column:amount;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;index;not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (e *Entry) TableName() string {
	return "entries"
}

func (e *Entry) ForeignKeys() []string {
	d := Diary{}
	ec := Category{}

	return []string{
		fmt.Sprintf("diary_id;%s(id);RESTRICT;RESTRICT", d.TableName()),
		fmt.Sprintf("category_id;%s(id);RESTRICT;RESTRICT", ec.TableName()),
	}
}
