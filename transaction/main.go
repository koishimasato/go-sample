package main

import (
	"fmt"
	"github.com/koishimasato/go-sample/transaction/application"
	"github.com/koishimasato/go-sample/transaction/domain/service"
	"github.com/koishimasato/go-sample/transaction/infrastructure/repository"
)

func main() {
	r := repository.NewUserRepository()
	s := service.NewUserService(r)
	a := application.NewUserApplicationService(r, s)

	err := a.Register("hoge")
	if err != nil {
		fmt.Println(err.Error())
	}
}
