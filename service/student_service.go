package service

import (
	"net/http"
	"sim-u/model"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type StudentService struct {
	StudentController model.StudentController
}

func NewStudentService(studentController model.StudentController) model.StudentService {
	return &StudentService{
		StudentController: studentController,
	}
}

func (s *StudentService) HandleRegisterUserAsStudent() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.RegisterRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return ErrInternal
		}

		register, err := s.StudentController.RegisterUserAsStudent(c.Request().Context(), model.RegisterRequest{
			Email:       req.Email,
			Name:        req.Name,
			PhoneNumber: req.PhoneNumber,
			Password:    req.Password,
			Repassword:  req.Repassword,
		})
		if err != nil {
			return ErrInternal
		}

		return c.JSON(http.StatusOK, register)
	}
}
