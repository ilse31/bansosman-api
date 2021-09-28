package routes

import (
	"bansosman/app/controllers/users"

	"github.com/labstack/echo/v4"
)

type HandlerRoute struct {
	UsersHandler users.ControllerUsers
}

func (handler *HandlerRoute) RouteRegister(e *echo.Echo) {
	user := e.Group("users")
	user.POST("/register", handler.UsersHandler.Register)
}
