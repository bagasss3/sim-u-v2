package model

import (
	"time"

	"gorm.io/gorm"
)

type EventStudent struct {
	Id        int64          `json:"id"`
	StudentId int64          `json:"student_id"`
	EventId   int64          `json:"event_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
