package router

import (
	"sim-u/model"

	"github.com/labstack/echo/v4"
)

type Service struct {
	group          *echo.Group
	studentService model.StudentService
}

func RouteService(group *echo.Group, studentService model.StudentService) {
	svc := &Service{
		group:          group,
		studentService: studentService,
	}
	svc.RouterInit()
}

func (s *Service) RouterInit() {
	s.routerStudent()
}
