package command

type UserGetCommand struct {
	ID string
}

type UserRegisterCommand struct {
	Name string
}

type UserUpdateCommand struct {
	ID string
	Name string
}

type UserDeleteCommand struct {
	ID string
}