package repository

import "collector/domain/entity"

// repo of channel
type IChannelRepository interface {
	Create(c *entity.Channel) (*entity.Channel, error)
	FindBySlackID(slackID string) (*entity.Channel, error)
	Update(c *entity.Channel) (*entity.Channel, error)
	Delete(c *entity.Channel) (*entity.Channel, error)
}
