package repository

import (
	"fmt"
	"github.com/koishimasato/go-sample/transaction/domain/model"
	"github.com/koishimasato/go-sample/transaction/domain/repository"
	"github.com/koishimasato/go-sample/transaction/infrastructure/db"
	"gopkg.in/gorp.v1"
)

type UserRepositoryImpl struct {
	Exec gorp.SqlExecutor
}

func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{
		Exec: db.GetDBMap(),
	}
}

func (r *UserRepositoryImpl) Find(user model.UserID) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryImpl) FindByName(name model.UserName) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryImpl) Save(user *model.User, tx *gorp.Transaction) error {
	err := tx.Insert(user)
	if err != nil {
		return err
	}

	fmt.Println("saved!")
	return nil
}

func (r *UserRepositoryImpl) Delete(user *model.User, tx *gorp.Transaction) error {
	fmt.Println("deleted!")
	return nil
}
