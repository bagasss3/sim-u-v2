package model

import (
	"time"

	"gorm.io/gorm"
)

type StudentOrganizationPrecondition struct {
	Id                    int64          `json:"id"`
	StudentOrganizationId int64          `json:"student_organization_id"`
	ImageId               int64          `json:"image_id"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
}
