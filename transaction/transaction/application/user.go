package application

import (
	"errors"
	"github.com/koishimasato/go-sample/transaction/domain/model"
	"github.com/koishimasato/go-sample/transaction/domain/repository"
	"github.com/koishimasato/go-sample/transaction/domain/service"
)

// UserApplicationService ユーザーアプリケーションサービス
type UserApplicationService struct {
	userRepository repository.UserRepository
	userService    *service.UserService
	uow repository.UnitOfWork
}

// NewUserApplicationService 新しいユーザーアプリケーションサービスを生成する
func NewUserApplicationService(r repository.UserRepository, s *service.UserService, uow repository.UnitOfWork) *UserApplicationService {
	return &UserApplicationService{userRepository: r, userService: s, uow: uow}
}

// Register ユーザーを登録する
func (s *UserApplicationService) Register(name string) error {
	// ユーザー名(値オブジェクト)を生成する
	userName, err := model.NewUserName(name)
	if err != nil {
		return err
	}
	// ユーザー(エンティティ)を生成する
	user, err := model.NewUser(nil, *userName)
	if err != nil {
		return err
	}

	// 存在を検証する
	exists, err := s.userService.Exists(user)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("ユーザーは既に存在しています")
	}

	// ユーザーを登録する (これはいらないのでは？)
	err = s.userRepository.Save(user)

	// uowにDB登録の仕事を丸投げする
	return s.uow.Commit()
}
