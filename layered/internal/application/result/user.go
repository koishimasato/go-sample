package result

import "github.com/koishimasato/go-sample/internal/domain/model"

type UserGetResult struct {
	user *model.User
}

func NewUserGetResult(user model.User) *UserGetResult {
	return &UserGetResult{user: &user}
}

type UserGetAllResult struct {
	users []*model.User
}

func NewUserGetAllResult(users []*model.User) *UserGetAllResult {
	return &UserGetAllResult{users: users}
}

type UserRegisterResult struct {
	user *model.User
}

func NewUserRegisterResult(user *model.User) *UserRegisterResult {
	return &UserRegisterResult{user: user}
}

