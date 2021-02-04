package repository

import (
	"github.com/koishimasato/go-sample/transaction/domain/repository"
	"gopkg.in/gorp.v1"
)

type UnitOfWorkImpl struct {
	transaction gorp.Transaction
	repository repository.UserRepository
}

func NewUnitOfWork() repository.UnitOfWork  {
	uow := &UnitOfWorkImpl{
		transaction: nil,
	}
	uow.repository = NewUserRepository(uow.transaction)
	return uow
}

func (u *UnitOfWorkImpl) Commit() error {
	err := u.transaction.Commit()
	return err
}

func (u *UnitOfWorkImpl) UserRepository() repository.UserRepository {
	return u.repository
}