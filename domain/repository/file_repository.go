package repository

import "collector/domain/entity"

type IFileRepository interface {
	Create(f *entity.File) (*entity.File, error)
	FindBySlackID(slackID string) (*entity.File, error)
	// FindByUser(user *entity.User) ([]*entity.File, error)
	Update(f *entity.File) (*entity.File, error)
	Delete(c *entity.File) (*entity.File, error)
}
