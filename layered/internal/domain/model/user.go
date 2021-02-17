package model

import (
	"errors"
)

type User struct {
	id   UserID
	name UserName
	userType UserType
}

func NewUser(id UserID, name UserName, userType UserType) (*User, error) {
	u := &User{id: id, name: name, userType: userType}
	return u, nil
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Name() UserName {
	return u.name
}

func (u *User) UserType() UserType {
	return u.userType
}

func (u *User) ChangeName(name UserName) {
	u.name = name
}

func (u *User) Upgrade() {
	u.userType = Premium
}

func (u *User) Downgrade() {
	u.userType = Normal
}

type UserID struct {
	value string
}

func NewUserID(val string) (*UserID, error) {
	if len(val) == 0 {
		return nil, errors.New("無効な値だよ")
	}
	return &UserID{value: val}, nil
}

func (u *UserID) Value() string {
	return u.value
}

type UserName struct {
	value string
}

func NewUserName(val string) (*UserName, error) {
	l := len(val)
	if l < 3 {
		return nil, errors.New("ユーザー名は3文字以上必要だよ")
	}
	if l > 20 {
		return nil, errors.New("ユーザー名は20文字までだよ")
	}
	return &UserName{value: val}, nil
}

func (u *UserName) Value() string {
	return u.value
}

type UserType string

const (
	Premium = UserType("Premium")
	Normal = UserType("Normal")
)

func (u *UserType) Value() string {
	return u.Value()
}

