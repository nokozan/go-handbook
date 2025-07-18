package main

import (
	"functional_options/service"
	"log"
	"time"
)

type customLogger struct{}

func (c customLogger) Info(msg string)  { println("[CUSTOM-INFO]", msg) }
func (c customLogger) Error(msg string) { println("[CUSTOM-ERROR]", msg) }

func main() {
	common := []service.Option{
		service.WithLogger(customLogger{}),
		service.WithProductionDefaults(),
		service.WithValidation(true),
		service.WithFeature("debug", true),
		service.WithFeature("cache", true),
	}

	svc, err := service.NewService(
		append(common,
			service.WithTimeout(5*time.Second),
		)...,
	)
	if err != nil {
		log.Fatal("Failed to create service : ", err)
	}

	svc.Run()

}
