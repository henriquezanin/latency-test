package service

import (
	"TCC2/service/handlers"
	"TCC2/service/handlers/models"
)

type Application struct {
	Routes handlers.Handler
}

func New(serviceChainURL string) Application {
	model := models.Params{ServiceChainURL: serviceChainURL}
	handler := handlers.New(model)
	return Application{Routes: handler}
}