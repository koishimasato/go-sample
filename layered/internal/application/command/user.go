package command

type UserGetCommand struct {
	ID string
}

func NewUserGetCommand(id string) UserGetCommand {
	return UserGetCommand{
		ID: id,
	}
}

type UserRegisterCommand struct {
	Name string
}

func NewUserRegisterCommand(name string) UserRegisterCommand {
	return UserRegisterCommand{
		Name: name,
	}
}

type UserUpdateCommand struct {
	ID   string
	Name string
}

func NewUserUpdateCommand(id string, name string) UserUpdateCommand {
	return UserUpdateCommand{
		ID:   id,
		Name: name,
	}
}

type UserDeleteCommand struct {
	ID string
}

func NewUserDeleteCommand(id string) UserDeleteCommand {
	return UserDeleteCommand{
		ID: id,
	}
}
