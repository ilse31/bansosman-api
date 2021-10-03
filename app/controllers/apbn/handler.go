package apbn

import (
	"bansosman/app/controllers/apbn/requset"
	"bansosman/app/controllers/apbn/response"
	"bansosman/bussiness/apbn"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ApbnController struct {
	serviceApbn apbn.Service
}

func NewHandler(apBnserv apbn.Service) *ApbnController {
	return &ApbnController{
		serviceApbn: apBnserv,
	}
}

func (handler *ApbnController) Create(echoConteks echo.Context) error {
	var req requset.ApbnReq

	if err := echoConteks.Bind(&req); err != nil {
		return echoConteks.JSON(http.StatusBadRequest, err)
	}
	domain := requset.ToDomain(req)

	resp, err := handler.serviceApbn.Append(domain)
	if err != nil {
		return echoConteks.JSON(http.StatusInternalServerError, err)
	}
	return echoConteks.JSON(http.StatusOK, response.FromDom(*resp))
}

func (handler *ApbnController) Update(echoOcnteks echo.Context) error {
	idstr := echoOcnteks.Param("id")
	id, err := strconv.Atoi(idstr)

	if err != nil {
		return echoOcnteks.JSON(http.StatusBadRequest, err)
	}
	var req requset.ApbnReq

	if err := echoOcnteks.Bind(&req); err != nil {
		return echoOcnteks.JSON(http.StatusBadRequest, err)
	}

	domain := requset.ToDomain(req)

	resp, err := handler.serviceApbn.Update(domain, id)

	if err != nil {
		return echoOcnteks.JSON(http.StatusInternalServerError, err)
	}
	return echoOcnteks.JSON(http.StatusOK, response.FromDom(*resp))
}

func (handler *ApbnController) ReadAll(echoconteks echo.Context) error {
	user, err := handler.serviceApbn.FindAll()
	if err != nil {
		return echoconteks.JSON(http.StatusBadRequest, err)
	}
	return echoconteks.JSON(http.StatusOK, response.NewResponseArray(user))
}

func (handler *ApbnController) ReadID(echoconteks echo.Context) error {
	idstr := echoconteks.Param("id")
	id, err := strconv.Atoi(idstr)

	if err != nil {
		return echoconteks.JSON(http.StatusBadRequest, err)
	}
	resp, err := handler.serviceApbn.FindByID(id)

	if err != nil {
		return echoconteks.JSON(http.StatusNotFound, err)
	}
	return echoconteks.JSON(http.StatusOK, response.FromDom(*resp))
}

func (handler *ApbnController) Delete(echoConteks echo.Context) error {
	idstr := echoConteks.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return echoConteks.JSON(http.StatusBadRequest, err)
	}
	user, err1 := handler.serviceApbn.FindByID(id)
	result, err2 := handler.serviceApbn.Delete(user, id)

	if err1 != nil {
		return echoConteks.JSON(http.StatusNotFound, err1)
	} else if err2 != nil {
		return echoConteks.JSON(http.StatusBadRequest, result)
	}

	return echoConteks.JSON(http.StatusOK, result)
}
