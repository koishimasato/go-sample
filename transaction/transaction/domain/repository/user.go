package repository

import (
	"github.com/koishimasato/go-sample/transaction/domain/model"
	_ "github.com/koishimasato/go-sample/transaction/domain/model"
	"gopkg.in/gorp.v1"
)

// UserRepository ユーザーリポジトリーインターフェース
type UserRepository interface {
	Find(model.UserID) (*model.User, error)
	FindByName(model.UserName) (*model.User, error)
	Save(*model.User, *gorp.Transaction) error
	Delete(*model.User, *gorp.Transaction) error
}

