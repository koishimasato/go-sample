package main

import (
	application "github.com/koishimasato/go-sample/internal/application/service"
	"github.com/koishimasato/go-sample/internal/domain/service"
	"github.com/koishimasato/go-sample/internal/infrastructure/factory"
	"github.com/koishimasato/go-sample/internal/infrastructure/repository"
	"github.com/koishimasato/go-sample/internal/presentation/handler"
	"net/http"
)

func main() {
	userFactory := factory.NewUserFactory()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userApplicationService := application.NewUserApplicationService(userFactory, userRepository, userService)
	userHandler := handler.NewUserHandler(userApplicationService)

	http.HandleFunc("/", userHandler.Handle)
	_ = http.ListenAndServe(":8080", nil)
}
