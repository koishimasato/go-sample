package service

import (
	"github.com/koishimasato/go-sample/internal/domain/model"
	"github.com/koishimasato/go-sample/internal/domain/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return UserService{userRepository: r}
}

func (s *UserService) Exists(user *model.User) (bool, error) {
	u, err := s.userRepository.FindByName(user.Name())
	if err != nil {
		return false, err
	}
	return u != nil, nil
}