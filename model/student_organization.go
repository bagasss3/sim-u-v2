package model

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type StudentOrganization struct {
	Id          int64          `json:"id"`
	UserId      int64          `json:"user_id"`
	ImageId     null.Int       `json:"image_id"`
	Description null.Int       `json:"description"`
	Status      StatusType     `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
