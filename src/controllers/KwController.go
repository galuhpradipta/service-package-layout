package controllers

import (
	"net/http"

	"github.com/galuhpradipta/service-package-layout/src/interfaces"
	"github.com/labstack/echo"
)

type KwController struct {
	interfaces.IKwService
}

func (ctrl *KwController) Welcome(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Welcome to API Service",
	})
}
