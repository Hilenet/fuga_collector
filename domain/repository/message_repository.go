package repository

import "collector/domain/entity"

type IMessageRepository interface {
	Create(m *entity.Message) (*entity.Message, error)
	FindByTs(ts string) (*entity.Message, error)
	// FindByUser(user *entity.User) ([]*entity.Message, error)
	// FindByChannel(user *entity.Channel) ([]*entity.Channel, error)
	Update(m *entity.Message) (*entity.Message, error)
	Delete(m *entity.Message) (*entity.Message, error)
}
