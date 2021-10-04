package routes

import (
	"bansosman/app/controllers/admin"
	"bansosman/app/controllers/apbn"
	"bansosman/app/controllers/daerah"
	"bansosman/app/controllers/users"

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

	//

	//Create Account
	creates := e.Group("create")
	//user
	creates.POST("/regist", handler.UsersHandler.Create)
	creates.GET("/log", handler.UsersHandler.Login)
	//admin
	creates.POST("/register", handler.AdminController.Register)
	creates.GET("/login", handler.AdminController.Login)

	//daerah
	daerah := e.Group("daerah")
	daerah.POST("/input", handler.Daerahhandler.Create)
	daerah.GET("/all", handler.Daerahhandler.ReadAll)
	daerah.GET("/:id", handler.Daerahhandler.ReadID)
	daerah.GET("/find-ip", handler.Daerahhandler.GetByIP)

	//auth val
	users := e.Group("user")
	users.DELETE("/:id", handler.UsersHandler.Delete)
	users.GET("/alluser", handler.UsersHandler.ReadAll)
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
