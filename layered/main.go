package main

import (
	"fmt"
	"github.com/koishimasato/go-sample/internal/application"
	"github.com/koishimasato/go-sample/internal/domain/service"
	"github.com/koishimasato/go-sample/internal/infrastructure/factory"
	"github.com/koishimasato/go-sample/internal/infrastructure/repository"
)

func main() {
	f := factory.NewUserFactory()
	r := repository.NewUserRepository()
	s := service.NewUserService(r)
	a := application.NewUserApplicationService(f, r, s)
	//fmt.Println(a.Get())
}
