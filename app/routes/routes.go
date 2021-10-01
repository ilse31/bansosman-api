package routes

import (
	"bansosman/app/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerRoute struct {
	UsersHandler  users.ControllerUsers
	JwtMiddleware middleware.JWTConfig
}

func (handler *HandlerRoute) RouteRegister(e *echo.Echo) {

	//crud users
	user := e.Group("users")
	user.POST("/register", handler.UsersHandler.Create)
	user.GET("/login", handler.UsersHandler.Login)
	//auth val
	users := e.Group("user", middleware.JWTWithConfig(handler.JwtMiddleware))
	users.DELETE("/:id", handler.UsersHandler.Delete)
	users.GET("/alluser", handler.UsersHandler.ReadAll)
	users.GET("/:id", handler.UsersHandler.ReadID)
	users.PUT("/updates", handler.UsersHandler.Update)

}
