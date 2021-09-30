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
	user.GET("/all", handler.UsersHandler.FindAll)
	user.GET("/:id", handler.UsersHandler.FindID)
	user.DELETE("/:id", handler.UsersHandler.Delete)
	user.PUT("/update", handler.UsersHandler.Update)
}
