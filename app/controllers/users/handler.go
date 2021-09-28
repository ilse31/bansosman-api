package users

import (
	"bansosman/app/controllers/users/request"
	"bansosman/app/controllers/users/response"
	"bansosman/bussiness/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ControllerUsers struct {
	serviceUser users.Service
}

func NewHandler(userServ users.Service) *ControllerUsers {
	return &ControllerUsers{
		serviceUser: userServ,
	}
}

func (handler *ControllerUsers) Register(echoConteks echo.Context) error {
	var req request.UsersRegist
	if err := echoConteks.Bind(&req); err != nil {
		return echoConteks.JSON(http.StatusBadRequest, err)
	}
	domain := request.ToDomain(req)
	resp, err := handler.serviceUser.Append(domain)
	if err != nil {
		return echoConteks.JSON(http.StatusInternalServerError, err)
	}
	return echoConteks.JSON(http.StatusOK, response.FromDomain(*resp))
}
