package config

import (
	"go-microservice-boilerplate/pkg/helper"
	"net/http"
)

func (s Endpoint) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	helper.JsonResponse(http.StatusOK, w, nil)
}
