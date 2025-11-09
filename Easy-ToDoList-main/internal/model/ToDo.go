package model

import (
	"gorm.io/gorm"
	"time"
)

type ToDo struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	UserId      uint      `json:"user_id"`
	EndTime     time.Time `json:"end_time"`
}

func (todo *ToDo) TableName() string {
	return "todo"
}
