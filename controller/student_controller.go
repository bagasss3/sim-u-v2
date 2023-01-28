package controller

import (
	"context"
	"sim-u/model"
)

type StudentController struct {
	StudentRepository model.StudentRepository
}

func NewStudentController(studentRepository model.StudentRepository) model.StudentController {
	return &StudentController{
		StudentRepository: studentRepository,
	}
}

func (s *StudentController) RegisterUserAsStudent(ctx context.Context, req model.RegisterRequest) (*model.Token, error) {
	return &model.Token{}, nil
}
