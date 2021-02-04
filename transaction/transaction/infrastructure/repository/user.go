package repository

import (
	"fmt"
	"github.com/koishimasato/go-sample/transaction/domain/model"
	"github.com/koishimasato/go-sample/transaction/domain/repository"
	"gopkg.in/gorp.v1"
)

type UserRepositoryImpl struct {
	tx gorp.Transaction
	users map[model.UserID]*model.User
}

func NewUserRepository(transaction gorp.Transaction) repository.UserRepository {
	return &UserRepositoryImpl{
		tx: transaction,
	}
}

func (r *UserRepositoryImpl) Find(userID model.UserID) (*model.User, error) {
	user := &model.User{} // user取得コード
	r.users[userID] = user
	
	return user, nil
}

func (r *UserRepositoryImpl) FindByName(name model.UserName) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryImpl) Save(user *model.User) error {
	// ユーザが存在する場合
	if _, ok := r.users[user.ID()]; ok {
		// Updateする
		_, err := r.tx.Update(user)
		if err != nil {
			return err
		}
	} else {
		// insertする
		err := r.tx.Insert(user)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *UserRepositoryImpl) Delete(user *model.User) error {
	fmt.Println("deleted!")
	return nil
}
