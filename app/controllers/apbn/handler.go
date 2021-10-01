package apbn

import (
	"bansosman/app/controllers/apbn/requset"
	"bansosman/app/controllers/apbn/response"
	"bansosman/bussiness/apbn"
	"net/http"

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
	var req requset.Apbn

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
