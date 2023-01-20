package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

func LogInfo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("Middleware")
		return next(c)
	}
}
