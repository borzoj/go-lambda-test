//+build wireinject

package main

import (
	httpClient "github.com/borzoj/go-lambda-test/pkg/http"
	weatherService "github.com/borzoj/go-lambda-test/pkg/weather"
	"github.com/google/wire"
)

func CreateWeatherService() (weatherService.Service, error) {
	wire.Build(
		httpClient.NewClient,
		weatherService.NewService,
	)
	return weatherService.Service{}, nil
}
