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
<<<<<<< HEAD
	user.GET("/alluser", handler.UsersHandler.ReadAll)
	user.GET("/:id", handler.UsersHandler.ReadID)
	user.PUT("/updates", handler.UsersHandler.Update)
=======
	user.GET("/all", handler.UsersHandler.FindAll)
	user.GET("/:id", handler.UsersHandler.FindID)
	user.DELETE("/:id", handler.UsersHandler.Delete)
	user.PUT("/update", handler.UsersHandler.Update)
>>>>>>> bb11f13760677954a5180cb5a72ff710b2d46945
}
