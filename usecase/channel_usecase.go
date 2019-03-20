package usecase

import (
	"collector/domain/entity"
	"collector/domain/repository"
)

// Interactorが集約してしまった
// Atomicを担保
// Input Dataはnlopes/slackに依存してる
type ChannelUsecase struct {
	repo repository.IChannelRepository
}

// initializer
func NewChannelUsecase(repo repository.IChannelRepository) ChannelUsecase {
	ChannelUsecase := ChannelUsecase{
		repo: repo,
	}
	return ChannelUsecase
}

// Slackから新規チャンネル
func (u *ChannelUsecase) AddChannel(channel *entity.Channel) (*entity.Channel, error) {
	res, err := u.repo.Create(channel)
	if nil != err {
		return nil, err
	}

	return res, nil
}

// Slackからチャンネル情報更新
func (u *ChannelUsecase) UpdateChannel(channel *entity.Channel) (*entity.Channel, error) {
	res, err := u.repo.Update(channel)
	if nil != err {
		return nil, err
	}

	return res, nil
}

// Slackからチャンネル削除
func (u *ChannelUsecase) DeleteChannel(channel *entity.Channel) (*entity.Channel, error) {
	res, err := u.repo.Update(channel)
	if nil != err {
		return nil, err
	}

	return res, nil
}
