package usecase

import (
	"collector/domain/entity"
	"collector/domain/repository"
)

type UserUsecase struct {
	repo repository.IUserRepository
}

// initializer
func NewUserUsecase(repo repository.IUserRepository) UserUsecase {
	userUsecase := UserUsecase{
		repo: repo,
	}
	return userUsecase
}

// Slackから新規ユーザ
func (u *UserUsecase) AddUser(user *entity.User) (*entity.User, error) {
	res, err := u.repo.Create(user)
	if nil != err {
		return nil, err
	}

	return res, nil
}

// Slackからユーザ情報更新
func (u *UserUsecase) UpdateUser(user *entity.User) (*entity.User, error) {
	res, err := u.repo.Update(user)
	if nil != err {
		return nil, err
	}

	return res, nil
}

// Slackからユーザ削除
func (u *UserUsecase) DeleteUser(user *entity.User) (*entity.User, error) {
	res, err := u.repo.Update(user)
	if nil != err {
		return nil, err
	}

	return res, nil
}
