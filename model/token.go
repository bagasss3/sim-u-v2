package model

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	Id        int64          `json:"id"`
	Email     string         `json:"email"`
	Token     string         `json:"token"`
	ExpiredAt time.Time      `json:"expired_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
