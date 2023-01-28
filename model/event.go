package model

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Event struct {
	Id                    int64          `json:"id"`
	StudentOrganizationId int64          `json:"student_organization_id"`
	Name                  string         `json:"name"`
	DateTime              time.Time      `json:"date_time"`
	ImageId               null.Int       `json:"image_id"`
	Description           string         `json:"description"`
	Precondition          int16          `json:"precondition"`
	Status                StatusType     `json:"status"`
	CategoryId            int64          `json:"category_id"`
	EligibilityId         int64          `json:"eligibility_id"`
	CreatedAt             time.Time      `json:"created_at"`
	CreatedBy             int64          `json:"created_by"`
	UpdatedAt             time.Time      `json:"updated_at"`
	UpdatedBy             int64          `json:"updated_by"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
	DeletedBy             null.Int       `json:"deleted_by"`
}
