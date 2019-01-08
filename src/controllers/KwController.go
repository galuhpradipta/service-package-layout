package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/galuhpradipta/service-package-layout/payloads"
	"github.com/galuhpradipta/service-package-layout/src/interfaces"
)

type KwController struct {
	interfaces.IKwService
}

func (c *KwController) Welcome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(payloads.Response{
		Code: http.StatusOK,
		Data: map[string]string{
			"message": "Welcome to API Service",
		},
	})
}
