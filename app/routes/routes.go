package routes

import (
	"bansosman/app/controllers"
	"bansosman/app/controllers/admin"
	"bansosman/app/controllers/apbn"
	"bansosman/app/controllers/users"
	m "bansosman/app/middleware"
	"bansosman/bussiness"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerRoute struct {
	UsersHandler  users.ControllerUsers
	AdminHandler  admin.AdminController
	Apbnhandler   apbn.ApbnController
	JwtMiddleware middleware.JWTConfig
}

func (handler *HandlerRoute) RouteRegister(e *echo.Echo) {

	//apbn
	apbns := e.Group("apbn")
	apbns.POST("/input", handler.Apbnhandler.Create)
	//Create Account
	user := e.Group("users")
	user.POST("/register", handler.UsersHandler.Create)
	user.GET("/login", handler.UsersHandler.Login)
	//
	admins := e.Group("admin")
	admins.POST("/reg", handler.AdminHandler.Register)
	admins.POST("/login", handler.AdminHandler.Login)

	//auth val
	users := e.Group("user", middleware.JWTWithConfig(handler.JwtMiddleware))
	users.DELETE("/:id", handler.UsersHandler.Delete)
	users.GET("/alluser", handler.UsersHandler.ReadAll)
	users.GET("/:id", handler.UsersHandler.ReadID)
	users.PUT("/updates", handler.UsersHandler.Update)

	//admin

}

func RoleValidation(roleID int, adminController admin.AdminController) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := m.GetUser(c)

			userRole := adminController.GetRoleByID(claims.ID)

			if userRole == roleID {
				return hf(c)
			} else {
				return controllers.NewErrorResponse(c, http.StatusForbidden, bussiness.ErrUnauthorized)
			}
		}
	}
}
