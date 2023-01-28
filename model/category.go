package model

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Category struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	Description null.String    `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	CreatedBy   int64          `json:"created_by"`
	UpdatedAt   time.Time      `json:"updated_at"`
	UpdatedBy   int64          `json:"updated_by"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	DeletedBy   null.Int       `json:"deleted_by"`
}
