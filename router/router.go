package router

import (
	"sim-u/controller"

	"github.com/labstack/echo/v4"
)

func RouterInit(router *echo.Echo) {
	router.GET("/", controller.Hello)
	router.GET("/helo", controller.Anjay)
}
