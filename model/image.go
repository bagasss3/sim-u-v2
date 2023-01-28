package model

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	Id         int64          `json:"id"`
	PublicId   string         `json:"public_id"`
	Width      int32          `json:"width"`
	Height     int32          `json:"height"`
	Version    int32          `json:"version"`
	Format     string         `json:"format"`
	Etag       string         `json:"etag"`
	Url        string         `json:"url"`
	Secure_url string         `json:"secure_url"`
	Signature  string         `json:"signature"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
