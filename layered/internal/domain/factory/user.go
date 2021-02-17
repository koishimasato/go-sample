package factory

import "github.com/koishimasato/go-sample/internal/domain/model"

// UserFactory ユーザーファクトリインターフェース
type UserFactory interface {
	Create(name model.UserName) (*model.User, error)
}

