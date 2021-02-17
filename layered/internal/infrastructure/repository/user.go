package repository

import (
	"github.com/koishimasato/go-sample/internal/domain/model"
	"github.com/koishimasato/go-sample/internal/domain/repository"
	"github.com/koishimasato/go-sample/internal/infrastructure/db"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Find(userID model.UserID) (*model.User, error) {
	var dbUser *db.User
	err := db.GetDBMap().SelectOne(&dbUser, "SELECT * FROM users WHERE id = ?", userID.Value())
	if err != nil {
		return nil, err
	}

	id, err := model.NewUserID(dbUser.ID)
	if err != nil {
		return nil, err
	}

	name, err := model.NewUserName(dbUser.Name)
	if err != nil {
		return nil, err
	}
	userType := model.NewUserType(dbUser.UserType)

	user, err := model.NewUser(*id, *name, userType)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindAll() ([]*model.User, error) {
	var dbUsers []*db.User
	var users []*model.User

	query := "SELECT * FROM users"
	_, err := db.GetDBMap().Select(&dbUsers, query)
	if err != nil {
		return nil, err
	}

	for _, user := range dbUsers {
		id, err := model.NewUserID(user.ID)
		if err != nil {
			return nil, err
		}

		name, err := model.NewUserName(user.Name)
		if err != nil {
			return nil, err
		}
		userType := model.NewUserType(user.UserType)

		u, err := model.NewUser(*id, *name, userType)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepositoryImpl) FindByName(name model.UserName) (*model.User, error) {
	return nil, nil
}

func (r *UserRepositoryImpl) Save(user *model.User) error {
	query := "INSERT INTO users (id, name, user_type) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE id=VALUES(id)"
	id := user.ID()
	name := user.Name()
	userType := user.UserType()
	_, err := db.GetDBMap().Exec(query, id.Value(), name.Value(), userType)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) Delete(user *model.User) error {
	id := user.ID()
	_, err := db.GetDBMap().Exec("DELETE FROM users WHERE id = ?", id.Value())
	if err != nil {
		return err
	}

	return nil
}
