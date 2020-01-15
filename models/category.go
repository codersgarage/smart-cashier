package models

import (
	"fmt"
	"time"
)

type Category struct {
	ID        string    `json:"id" gorm:"column:id;unique_index;not null"`
	DiaryID   string    `json:"-" gorm:"column:diary_id;primary_key;not null"`
	Name      string    `json:"name" gorm:"column:name;primary_key;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null;index"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type CategoryDetails struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

func (ec *Category) TableName() string {
	return "categories"
}

func (ec *Category) ForeignKeys() []string {
	d := Diary{}

	return []string{
		fmt.Sprintf("diary_id;%s(id);RESTRICT;RESTRICT", d.TableName()),
	}
}
