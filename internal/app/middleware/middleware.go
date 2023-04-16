package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		role := ctx.Request().Header.Get("User-Role")

		if role == "admin" {
			log.Println("Admin role detected!")
		}

		err := next(ctx)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	}
}
