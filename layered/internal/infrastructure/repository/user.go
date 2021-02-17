package repository

import (
	"fmt"
	"github.com/koishimasato/go-sample/internal/domain/model"
	"github.com/koishimasato/go-sample/internal/domain/repository"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Find(userID model.UserID) (*model.User, error) {
	name, err := model.NewUserName("hoge") // TODO: DBからの検索コードが入る
	if err != nil {
		return nil, err
	}
	user, err := model.NewUser(userID, *name, model.Normal)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindAll() ([]*model.User, error) {
	return []*model.User{}, nil
}

func (r *UserRepositoryImpl) FindByName(name model.UserName) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryImpl) Save(user *model.User) error {
	return nil
}

func (r *UserRepositoryImpl) Delete(user *model.User) error {
	fmt.Println("deleted!")
	return nil
}
