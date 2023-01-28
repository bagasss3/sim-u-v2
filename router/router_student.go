package router

func (s *Service) routerStudent() {
	s.group.GET("/participant", s.studentService.HandleRegisterUserAsStudent())
}
