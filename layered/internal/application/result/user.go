package result

import "github.com/koishimasato/go-sample/internal/domain/model"

type UserGetResult struct {
	User *model.User
}

func NewUserGetResult(user model.User) *UserGetResult {
	return &UserGetResult{User: &user}
}

type UserGetAllResult struct {
	Users []*model.User
}

func NewUserGetAllResult(users []*model.User) *UserGetAllResult {
	return &UserGetAllResult{Users: users}
}

type UserRegisterResult struct {
	User *model.User
}

func NewUserRegisterResult(user *model.User) *UserRegisterResult {
	return &UserRegisterResult{User: user}
}

