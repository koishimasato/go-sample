package repositories

import (
	"fmt"
	"github.com/koishimasato/go-sample/transaction/domain/model"
	"github.com/koishimasato/go-sample/transaction/domain/repository"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Find(user model.UserID) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryImpl) FindByName(name model.UserName) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryImpl) Save(user *model.User) error {
	fmt.Println("saved!")
	return nil
}

func (r *UserRepositoryImpl) Delete(user *model.User) error {
	fmt.Println("deleted!")
	return nil
}
