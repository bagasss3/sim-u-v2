package model

import (
	"time"

	"gorm.io/gorm"
)

type Role string
type StatusType string

const (
	RoleStudent             Role = "STUDENT"
	RoleStudentOrganization Role = "STUDENT_ORGANIZATION"
	RoleAdmin               Role = "ADMIN"
)

const (
	StatusPending   StatusType = "PENDING"
	StatusActive    StatusType = "ACTIVE"
	StatusNonActive StatusType = "NONACTIVE"
)

type User struct {
	Id        int64          `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Role      Role           `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
