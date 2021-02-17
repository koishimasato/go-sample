package service

import (
	"database/sql"
	"errors"
	"github.com/koishimasato/go-sample/internal/application/command"
	"github.com/koishimasato/go-sample/internal/application/result"
	"github.com/koishimasato/go-sample/internal/domain/factory"
	"github.com/koishimasato/go-sample/internal/domain/model"
	"github.com/koishimasato/go-sample/internal/domain/repository"
	"github.com/koishimasato/go-sample/internal/domain/service"
)

// UserApplicationService ユーザーアプリケーションサービス
type UserApplicationService struct {
	userFactory    factory.UserFactory
	userRepository repository.UserRepository
	userService    service.UserService
}

func NewUserApplicationService(f factory.UserFactory, r repository.UserRepository, s service.UserService) UserApplicationService {
	return UserApplicationService{userFactory: f, userRepository: r, userService: s}
}

func (s *UserApplicationService) Get(command command.UserGetCommand) (*result.UserGetResult, error) {
	id, err := model.NewUserID(command.ID)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepository.Find(*id)
	if err != nil {
		return nil, err
	}
	return result.NewUserGetResult(*user), err
}

func (s *UserApplicationService) GetAll() (*result.UserGetAllResult, error) {
	users, err := s.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return result.NewUserGetAllResult(users), err
}

func (s *UserApplicationService) Register(command command.UserRegisterCommand) (*result.UserRegisterResult, error) {
	userName, err := model.NewUserName(command.Name)
	if err != nil {
		return nil, err
	}
	user, err := s.userFactory.Create(*userName)
	if err != nil {
		return nil, err
	}

	exists, err := s.userService.Exists(user)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("ユーザーは既に存在しています")
	}

	err = s.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return result.NewUserRegisterResult(user), nil
}

func (s *UserApplicationService) Update(command command.UserUpdateCommand) error {
	id, err := model.NewUserID(command.ID)
	if err != nil {
		return err
	}
	user, err := s.userRepository.Find(*id)
	if user == nil || err == sql.ErrNoRows {
		return errors.New("ユーザーが存在しません")
	}
	if err != nil {
		return err
	}

	name, err := model.NewUserName(command.Name)
	if err != nil {
		return err
	}
	user.ChangeName(*name)

	exists, err := s.userService.Exists(user)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("ユーザーは既に存在しています")
	}

	err = s.userRepository.Save(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserApplicationService) Delete(command command.UserDeleteCommand) error {
	id, err := model.NewUserID(command.ID)
	if err != nil {
		return err
	}
	user, err := s.userRepository.Find(*id)
	if err != nil {
		return err
	}

	if user == nil {
		return nil
	}

	err = s.userRepository.Delete(user)
	if err != nil {
		return err
	}

	return nil
}
