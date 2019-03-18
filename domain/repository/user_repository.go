package repository

import "collector/domain/entity"

type IUserRepository interface {
	Create(u *entity.User) (*entity.User, error)
	FindbySlackID(slackID string) (*entity.User, error)
	// FindbyChannel(Channel string) (*entity.User, error)
	Update(u *entity.User) (*entity.User, error)
	Delete(u *entity.User) (*entity.User, error)
}
