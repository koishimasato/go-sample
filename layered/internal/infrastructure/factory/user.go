package factory

import (
	"github.com/google/uuid"
	"github.com/koishimasato/go-sample/internal/domain/factory"
	"github.com/koishimasato/go-sample/internal/domain/model"
	"log"
)

type UserFactoryImpl struct {
}

func NewUserFactory() factory.UserFactory {
	return &UserFactoryImpl{
	}
}

func (r *UserFactoryImpl) Create(name model.UserName) (*model.User, error) {
	newUUID, err := uuid.NewUUID()
	log.Println(newUUID)
	if err != nil {
		return nil, err
	}
	id, err := model.NewUserID(newUUID.String())
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	user, err := model.NewUser(*id, name, model.Normal) // user取得コード
	if err != nil {
		return nil, err
	}

	return user, nil
}
