package repository

type UnitOfWork interface {
	Commit() error
	UserRepository() UserRepository
}
