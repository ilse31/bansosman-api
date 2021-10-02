package daerah

import (
	"bansosman/app/controllers/daerah/request"
	"bansosman/app/controllers/daerah/response"
	"bansosman/bussiness/daerah"
	"net/http"

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
