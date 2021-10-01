package admin

import (
	"bansosman/app/controllers/admin/request"
	"bansosman/bussiness/admin"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	adminServ admin.Service
}

func NewUserController(service admin.Service) *AdminController {
	return &AdminController{
		adminServ: service,
	}
}

func (handler *AdminController) Register(c echo.Context) error {
	req := request.Admins{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data, err := handler.adminServ.Register(req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, data)
}

func (handler *AdminController) Login(c echo.Context) error {
	req := request.AdminLogin{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, err := handler.adminServ.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}

	res := struct {
		Token string `json:"token"`
	}{Token: token}

	return c.JSON(http.StatusOK, res)
}

func (handler *AdminController) GetRoleByID(id int) int {
	user, err := handler.adminServ.GetByID(uint(id))
	if err != nil {
		return -1
	}
	return int(user.RoleID)
}
