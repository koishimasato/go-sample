package main

import (
	"fmt"
	"github.com/koishimasato/go-sample/transaction/application"
	"github.com/koishimasato/go-sample/transaction/domain/service"
	"github.com/koishimasato/go-sample/transaction/infrastructure/repositories"
)

func main() {
	r := repositories.NewUserRepository()
	s := service.NewUserService(r)
	a := application.NewUserApplicationService(r, s)

	err := a.Register("hoge")
	if err != nil {
		fmt.Println("登録に失敗しました。")
	}
}
