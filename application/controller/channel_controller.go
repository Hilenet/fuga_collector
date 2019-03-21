package controller

import (
	"collector/domain/entity"
	"collector/domain/repository"
	"collector/usecase"

	"github.com/nlopes/slack"
)

// 上位層で発生した外部イベントを内部イベントに変換
// この際データの整形も担当，InputDataはentityで
type ChannelController struct {
	Usecase usecase.IChannelUsecase
}

// initializer
func NewChannelController(repoCollection *repository.RepositoryCollection) ChannelController {
	usecase := usecase.NewChannelUsecase(repoCollection.ChannelRepository, repoCollection.UserRepository)
	ChannelController := ChannelController{
		Usecase: &usecase,
	}
	return ChannelController
}

// slack.ChannelCreatedEvent -> AddChannel
// 外部参照の解決はStore時で良い？entityを生成する必要も特になさげ
/*
func (c *ChannelController) HandleChannelCreated(info *slack.ChannelCreatedInfo) error {
	// DMの場合は弾く
	if !info.IsChannel {
		return nil
	}

	// tmp, authorはIDで引回すか
	channel, err := entity.NewChannel(info.ID, info.Name, new(entity.User), entity.Topic{}, entity.Purpose{})
	if err != nil {
		return err
	}
	_, err = c.Usecase.AddChannel(&channel)
	if err != nil {
		return err
	}

	return nil
}
*/

// slack.ChannelJoinedEvent -> UpdateChannel
func (c *ChannelController) HandleChannelJoined(e *slack.ChannelJoinedEvent) error {
	user := c.Usecase.kkk
	_, err := c.Usecase.UpdateChannel(e.Channel.ID, e.Channel.Name, e.Channel.Creator)
	if err != nil {
		return err
	}

	return nil
}

// ChannelChanged
func (c *ChannelController) HandleChannelChange(Channel entity.Channel) error {
	_, err := c.Usecase.UpdateChannel(&Channel)
	if err != nil {
		return err
	}

	return nil
}

func (c *ChannelController) HandleChannelDelete(Channel entity.Channel) error {
	_, err := c.Usecase.DeleteChannel(&Channel)
	if err != nil {
		return err
	}

	return nil
}
