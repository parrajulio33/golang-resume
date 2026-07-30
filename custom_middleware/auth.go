package custom_middleware

import (
	"net/http"

	"resume-app/session" // adjust import path

	"github.com/labstack/echo/v4"
)

func RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if session.GetUserID(c) == "" {
			return c.Redirect(http.StatusFound, "/admin/")
		}

		c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, private")
		c.Response().Header().Set("Pragma", "no-cache")
		c.Response().Header().Set("Expires", "0")

		return next(c)
	}
}
