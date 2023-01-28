package repository

import (
	"context"
	"sim-u/model"

	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) model.StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (s *StudentRepository) Store(ctx context.Context) {

}

func (s *StudentRepository) FindById(ctx context.Context) {

}
