package usecase

import (
	"collector/domain/entity"
	"collector/domain/repository"
)

// Interactorが集約してしまった
// Atomicを担保
// Input Dataはnlopes/slackに依存してる
type ChannelUsecase struct {
	ChannelRepository repository.IChannelRepository
	UserRepository    repository.IUserRepository
}

// initializer
func NewChannelUsecase(channelRepo repository.IChannelRepository, userRepo repository.IUserRepository) ChannelUsecase {
	ChannelUsecase := ChannelUsecase{
		ChannelRepository: channelRepo,
		UserRepository:    userRepo,
	}
	return ChannelUsecase
}

// Slackから新規チャンネル
func (u *ChannelUsecase) AddChannel(channel *entity.Channel) (*entity.Channel, error) {
	res, err := u.ChannelRepository.Create(channel)
	if nil != err {
		return nil, err
	}

	return res, nil
}

// Slackからチャンネル情報更新
func (u *ChannelUsecase) UpdateChannel(channel *entity.Channel) (*entity.Channel, error) {
	res, err := u.ChannelRepository.Update(channel)
	if nil != err {
		return nil, err
	}

	return res, nil
}

// Slackからチャンネル削除
func (u *ChannelUsecase) DeleteChannel(channel *entity.Channel) (*entity.Channel, error) {
	res, err := u.ChannelRepository.Update(channel)
	if nil != err {
		return nil, err
	}

	return res, nil
}
