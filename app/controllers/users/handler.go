package users

import (
	"bansosman/app/controllers/users/request"
	"bansosman/app/controllers/users/response"
	"bansosman/bussiness/users"
	"net/http"
	"strconv"

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

<<<<<<< HEAD
func (handler *ControllerUsers) Create(echoContext echo.Context) error {
	var req request.UsersRegist
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, err)
=======
func (handler *ControllerUsers) Create(echoConteks echo.Context) error {
	var req request.UsersRegist
	if err := echoConteks.Bind(&req); err != nil {
		return echoConteks.JSON(http.StatusBadRequest, err)
>>>>>>> bb11f13760677954a5180cb5a72ff710b2d46945
	}
	domain := request.ToDomain(req)
	resp, err := handler.serviceUser.Append(domain)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, err)
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *ControllerUsers) Update(echoContext echo.Context) error {
	var req request.UsersUpdate
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, err)
	}
	domain := request.ToDomainUpdate(req)
	resp, err := handler.serviceUser.Update(domain)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, err)
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *ControllerUsers) ReadAll(echoContext echo.Context) error {
	user, err := handler.serviceUser.FindAll()
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, err)
	}
	return echoContext.JSON(http.StatusOK, response.NewResponseArray(user))
}

func (handler *ControllerUsers) ReadID(echoContext echo.Context) error {
	idstr := echoContext.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, err)
	}
	resp, err := handler.serviceUser.FindByID(id)
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, err)
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *ControllerUsers) Delete(echoContext echo.Context) error {
	idstr := echoContext.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, err)
	}
	game, err1 := handler.serviceUser.FindByID(id)
	result, err2 := handler.serviceUser.Delete(game, id)

	if err1 != nil {
		return echoContext.JSON(http.StatusNotFound, err1)
	} else if err2 != nil {
		return echoContext.JSON(http.StatusBadRequest, err2)
	}
	return echoConteks.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *ControllerUsers) FindAll(echoConteks echo.Context) error {
	users, err := handler.serviceUser.FindAll()
	if err != nil {
		return echoConteks.JSON(http.StatusBadRequest, err)
	}
	return echoConteks.JSON(http.StatusOK, response.NewResponseArray(users))
}

func (handler *ControllerUsers) FindID(echoConteks echo.Context) error {
	id := echoConteks.Param("id")
	idc, err := strconv.Atoi(id)
	if err != nil {
		return echoConteks.JSON(http.StatusBadRequest, err)
	}
	respon, err := handler.serviceUser.FindByID(idc)
	if err != nil {
		return echoConteks.JSON(http.StatusBadRequest, err)
	}
	return echoConteks.JSON(http.StatusOK, response.FromDomain(*respon))
}

<<<<<<< HEAD
	return echoContext.JSON(http.StatusOK, result)
=======
func (handler *ControllerUsers) Delete(echoConteks echo.Context) error {
	idstr := echoConteks.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return echoConteks.JSON(http.StatusBadRequest, err)
	}
	user, err1 := handler.serviceUser.FindByID(id)
	result, err2 := handler.serviceUser.Delete(user, id)

	if err1 != nil {
		return echoConteks.JSON(http.StatusNotFound, err1)
	} else if err2 != nil {
		return echoConteks.JSON(http.StatusBadRequest, err2)
	}

	return echoConteks.JSON(http.StatusOK, result)
}

func (handler *ControllerUsers) Update(echoConteks echo.Context) error {
	var req request.UsersUpdate
	if err := echoConteks.Bind(&req); err != nil {
		return echoConteks.JSON(http.StatusBadRequest, err)
	}
	domain := request.ToDomainupd(req)
	resp, err := handler.serviceUser.Update(domain)
	if err != nil {
		return echoConteks.JSON(http.StatusInternalServerError, err)
	}
	return echoConteks.JSON(http.StatusOK, response.FromDomain(*resp))
>>>>>>> bb11f13760677954a5180cb5a72ff710b2d46945
}
