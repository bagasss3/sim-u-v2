package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "MALE"
	Female Gender = "FEMALE"
)

type RegisterRequest struct {
	Email       string
	Name        string
	PhoneNumber string
	Password    string
	Repassword  string
}

type Student struct {
	Id          int64          `json:"id"`
	UserId      int64          `json:"user_id"`
	PhoneNumber string         `json:"phone_number"`
	Gender      *Gender        `json:"gender"`
	BirthDate   null.Time      `json:"birth_date"`
	ImageId     null.Int       `json:"image_id"`
	Status      StatusType     `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type StudentService interface {
	HandleRegisterUserAsStudent() echo.HandlerFunc
}

type StudentController interface {
	RegisterUserAsStudent(ctx context.Context, req RegisterRequest) (*Token, error)
}

type StudentRepository interface {
	Store(ctx context.Context)
	FindById(ctx context.Context)
}
