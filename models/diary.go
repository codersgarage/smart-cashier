package models

import (
	"fmt"
	"time"
)

type Diary struct {
	ID        string    `json:"id" gorm:"column:id;unique_index;not null"`
	UserID    string    `json:"-" gorm:"column:user_id;primary_key;not null"`
	Name      string    `json:"name" gorm:"column:name;primary_key;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;index;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (d *Diary) TableName() string {
	return "diaries"
}

func (d *Diary) ForeignKeys() []string {
	u := User{}

	return []string{
		fmt.Sprintf("user_id;%s(id);RESTRICT;RESTRICT", u.TableName()),
	}
}
