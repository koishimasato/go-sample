package model

import (
	"errors"
	"github.com/google/uuid"
)

// User ユーザー型
type User struct {
	id   UserID
	name UserName
}

// NewUser ユーザーを生成する
func NewUser(id *UserID, name UserName) (*User, error) {
	if id == nil {
		uuid, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		id, err = NewUserID(uuid.String())
		if err != nil {
			return nil, err
		}
	}
	return &User{id: *id, name: name}, nil
}

// ID ユーザーのIDを取得する
func (u *User) ID() UserID {
	return u.id
}

// Name ユーザーの名前を取得する
func (u *User) Name() UserName {
	return u.name
}

// ChangeName ユーザーの名前を変更する
func (u *User) ChangeName(name UserName) {
	u.name = name
}


// UserID ユーザーID型
type UserID struct {
	value string
}

// NewUserID 新規ユーザーIDを生成する
func NewUserID(val string) (*UserID, error) {
	if len(val) == 0 {
		return nil, errors.New("無効な値だよ")
	}
	return &UserID{value: val}, nil
}

// Value 値を取得する
func (u *UserID) Value() string {
	return u.value
}

// UserName ユーザー名型
type UserName struct {
	value string
}

// NewUserName ユーザー名を生成する
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

// Value 値を取得する
func (u *UserName) Value() string {
	return u.value
}