package factory

import "github.com/koishimasato/go-sample/internal/domain/model"

type UserFactory interface {
	Create(name model.UserName) (*model.User, error)
}

