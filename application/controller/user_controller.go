package controller

import (
	"collector/domain/entity"
	"collector/domain/repository"
	"collector/usecase"
)

// DAOをいい感じにしてusecaseに投げる
type UserController struct {
	Usecase usecase.IUserUsecase
}

// initializer
func NewUserController(repo repository.IUserRepository) UserController {
	usecase := usecase.NewUserUsecase(repo)
	userController := UserController{
		Usecase: &usecase,
	}
	return userController
}

// TeamJoin
func (c *UserController) HandleTeamJoin(user entity.User) error {
	_, err := c.Usecase.AddUser(&user)
	if err != nil {
		return err
	}

	return nil
}

// UserChange
func (c *UserController) HandleUserChange(user entity.User) error {
	_, err := c.Usecase.UpdateUser(&user)
	if err != nil {
		return err
	}

	return nil
}
