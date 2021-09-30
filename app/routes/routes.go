package routes

import (
	"bansosman/app/controllers/users"

	"github.com/labstack/echo/v4"
)

type HandlerRoute struct {
	UsersHandler users.ControllerUsers
}

func (handler *HandlerRoute) RouteRegister(e *echo.Echo) {

	//crud users
	user := e.Group("users")
	user.POST("/register", handler.UsersHandler.Create)
	user.GET("/alluser", handler.UsersHandler.ReadAll)
	user.GET("/:id", handler.UsersHandler.ReadID)
	user.PUT("/updates", handler.UsersHandler.Update)
}
