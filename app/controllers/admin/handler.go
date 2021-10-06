package admin

import (
	"bansosman/app/controllers"
	"bansosman/app/controllers/admin/request"
	"bansosman/bussiness/admin"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService admin.Service
}

func NewUserController(service admin.Service) *UserController {
	return &UserController{
		userService: service,
	}
}

func (ctrl *UserController) Register(c echo.Context) error {
	req := request.Admins{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.userService.Register(req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, data)
}

func (ctrl *UserController) Login(c echo.Context) error {

	req := request.Admins{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.userService.Login(req.Username, req.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	res := struct {
		Token string `json:"token"`
	}{Token: token}

	return controllers.NewSuccessResponse(c, res)
}

func (ctrl *UserController) GetRoleByID(id int) int {
	user, err := ctrl.userService.GetByID(uint(id))
	if err != nil {
		return -1
	}
	return int(user.RoleID)
}
