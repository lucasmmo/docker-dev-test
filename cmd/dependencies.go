package main

import "go-microservice-boilerplate/pkg/api/service/config"

type dependencies struct {
	configEndpoint *config.Endpoint
}

func InitDependencies() *dependencies {

	config_endpoint := config.NewEndpoint()

	return &dependencies{
		configEndpoint: config_endpoint,
	}
}
