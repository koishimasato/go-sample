package service

import (
	"github.com/koishimasato/go-sample/internal/domain/model"
	"github.com/koishimasato/go-sample/internal/domain/repository"
)

// UserService ユーザーサービス
type UserService struct {
	userRepository repository.UserRepository
}

// NewUserService 新しいユーザーサービスを生成する
func NewUserService(r repository.UserRepository) UserService {
	return UserService{userRepository: r}
}

// Exists 指定したユーザーが存在するか否か
func (s *UserService) Exists(user *model.User) (bool, error) {
	u, err := s.userRepository.FindByName(user.Name())
	if err != nil {
		return false, err
	}
	return u != nil, nil
}