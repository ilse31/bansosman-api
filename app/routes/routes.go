package routes

import (
	"bansosman/app/controllers/admin"
	"bansosman/app/controllers/apbn"
	"bansosman/app/controllers/daerah"
	"bansosman/app/controllers/users"
	m "bansosman/app/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerRoute struct {
	UsersHandler    users.ControllerUsers
	AdminController admin.UserController
	Apbnhandler     apbn.ApbnController
	Daerahhandler   daerah.Daerahcontroller
	JwtMiddleware   middleware.JWTConfig
}

func (handler *HandlerRoute) RouteRegister(e *echo.Echo) {

	//Create Account
	creates := e.Group("create")
	//user
	creates.POST("/register", handler.UsersHandler.Create)
	creates.GET("/login", handler.UsersHandler.Login)
	//admin
	creates.POST("/register", handler.AdminController.Register)
	creates.GET("/login", handler.AdminController.Login)

	//daerah
	daerah := e.Group("daerah")
	daerah.POST("/input", handler.Daerahhandler.Create)
	daerah.GET("/all", handler.Daerahhandler.ReadAll)
	daerah.GET("/:id", handler.Daerahhandler.ReadID)

	//auth val
	users := e.Group("user", middleware.JWTWithConfig(handler.JwtMiddleware))
	users.DELETE("/:id", handler.UsersHandler.Delete)
	users.GET("/alluser", handler.UsersHandler.ReadAll, m.RoleValidation("admin"))
	users.GET("/:id", handler.UsersHandler.ReadID)
	users.PUT("/updates/:id", handler.UsersHandler.Update)

	//apbn
	apbns := e.Group("apbn")
	apbns.POST("/input", handler.Apbnhandler.Create)
	apbns.GET("/all", handler.Apbnhandler.ReadAll)
	apbns.GET("/:id", handler.Apbnhandler.ReadID)
	apbns.PUT("/update/:id", handler.Apbnhandler.Update)
	apbns.DELETE("/:id", handler.Apbnhandler.Delete)
}
