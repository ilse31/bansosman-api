package middleware

import (
	"bansosman/app/controllers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RoleValidation(role string) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := GetUser(c)

			if claims.Role == role {
				return hf(c)
			} else {
				return controllers.NewErrorResponse(c, http.StatusForbidden, errors.New("forbidden roles"))
			}
		}
	}
}
