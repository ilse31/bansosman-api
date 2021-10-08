package daerah

import (
	"bansosman/app/controllers"
	"bansosman/app/controllers/daerah/request"
	"bansosman/app/controllers/daerah/response"
	"bansosman/bussiness/daerah"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Daerahcontroller struct {
	service daerah.Service
}

func Newhandler(daerahServ daerah.Service) *Daerahcontroller {
	return &Daerahcontroller{
		service: daerahServ,
	}
}

func (handler *Daerahcontroller) Create(echoConteks echo.Context) error {
	var req request.Daerah

	if err := echoConteks.Bind(&req); err != nil {
		return echoConteks.JSON(http.StatusBadRequest, err)
	}
	domain := request.ToDomain(req)

	resp, err := handler.service.Append(domain)
	if err != nil {
		return echoConteks.JSON(http.StatusInternalServerError, err)
	}
	return echoConteks.JSON(http.StatusOK, response.FromDom(*resp))
}

func (handler *Daerahcontroller) ReadAll(echoconteks echo.Context) error {
	user, err := handler.service.FindAll()
	if err != nil {
		return echoconteks.JSON(http.StatusBadRequest, err)
	}
	return echoconteks.JSON(http.StatusOK, response.NewResponseArray(user))
}

func (handler *Daerahcontroller) ReadID(echoconteks echo.Context) error {
	idstr := echoconteks.Param("id")
	id, err := strconv.Atoi(idstr)

	if err != nil {
		return echoconteks.JSON(http.StatusBadRequest, err)
	}
	resp, err := handler.service.FindByID(id)

	if err != nil {
		return echoconteks.JSON(http.StatusNotFound, err)
	}
	return echoconteks.JSON(http.StatusOK, response.FromDom(*resp))
}

func (handler *Daerahcontroller) GetByIP(c echo.Context) error {
	data, err := handler.service.GetByIP()
	if len(data) == 0 {
		return controllers.NewErrorResponse(c, http.StatusNotFound, err)
	}
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, data)
}
