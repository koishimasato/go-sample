package repository

import (
	"github.com/koishimasato/go-sample/internal/domain/model"
	_ "github.com/koishimasato/go-sample/internal/domain/model"
)

// UserRepository ユーザーリポジトリーインターフェース
type UserRepository interface {
	Find(model.UserID) (*model.User, error)
	FindByName(model.UserName) (*model.User, error)
	FindAll() ([]*model.User, error)
	Save(*model.User) error
	Delete(*model.User) error
}

