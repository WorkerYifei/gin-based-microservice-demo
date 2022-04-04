// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"gin-based-microservice-demo/internal/application"
)

// Injectors from wire.go:

func initializeApplication(name application.Name) (*application.Application, error) {
	applicationApplication := application.New(name)
	return applicationApplication, nil
}
