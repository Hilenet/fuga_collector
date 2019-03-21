package handler

import (
	"collector/application/controller"
	"collector/domain/entity"
	"collector/domain/repository"

	"fmt"

	"github.com/nlopes/slack"
)

// streamバラしてgoroutineに投げる
func HandleStream(slack *SlackHandler, repos *repository.RepositoryCollection) {
	controllers := controller.NewControllerCollection(repos)

	for msg := range slack.RTM.IncomingEvents {
		go handleEvent(&controllers, msg)
	}
}

// 変換して，application層に受け渡し
func handleEvent(conts *controller.ControllerCollection, msg *slack.RTMEvent) {
	switch ev := msg.Data.(type) {
	case *slack.HelloEvent:
		fmt.Println("hello")
	// channnel
	case *slack.ChannelCreatedEvent, *slack.ChannelDeletedEvent,
		*slack.ChannelArchivedEvent, *slack.ChannelRenameEvent:
		handleChannelEvent(conts, ev)
	// message
	case *slack.MessageEvent:
		// handleMessageEvent(&conts.MessageController, ev)
		handleMessageEvent(conts, ev)
	// user系
	case *slack.TeamJoinEvent, *slack.UserChangeEvent,
		*slack.ChannelJoinedEvent, *slack.ChannelLeftEvent:
		handleUserEvent(&conts.UserController, msg)
	default:
		// fmt.Printf("%v\n", ev)
		fmt.Printf("\n")
	}
}

// channelに係るEventを適切にhandle
// TODO: 汚いのでmappingをconfigでrouter的に実装したいね
func handleChannelEvent(conts *controller.ControllerCollection, ev *slack.RTMEvent) {
	cc := conts.ChannelController
	uc := conts.UserController
	// golangはfall throughではない
	switch e := ev.Data.(type) {
	case *slack.ChannelCreatedEvent:
		c.HandleChannelCreated(e)
	case *slack.ChannelJoinedEvent:
		c.HandleChannelJoined(e)
	case *slack.ChannelLeftEvent:
		c.HandleChannelLeft(e)
	case *slack.ChannelDeletedEvent:
		c.HandleChannelDeleted(e)
	case *slack.ChannelInfoEvent:
		c.HandleChannelInfo(e)
	case *slack.ChannelRenameEvent:
		c.HandleChannelRename(e)
	}
}

// func handleMessageEvent(cc *controller.ControllerCollection, ev *slack.MessageEvent) {
func handleMessageEvent(c *controller.MessageController, ev *slack.MessageEvent) {
	// channel, err := cc.ChannelController.Usecase.FindChannel(ev.Channel)
	// if err != nil {
	// 	return
	// }

	// user, err := cc.UserController.Usecase.FindUser(ev.User)
	// if err != nil {
	// 	return
	// }

	// 修正して受け取る，Fileの配列makeしてるけどちょっと要検証
	// message, err := entity.NewMessage(ev.Timestamp, channel, user, ev.Text, make([]string, 0))
	// mc := cc.MessageController

	switch ev.SubType {
	case "":
		c.HandleNewMessage(message)

	case "message_changed":
		c.HandleMessageChange(message)

	case "message_deleted":
		c.HandleMessageDelete(message)

		// case "message_replied":
		// default:
	}
}

func handleUserEvent(c *controller.UserController, msg slack.RTMEvent) {
	// controllerに投げる
	switch ev := msg.Data.(type) {
	case *slack.TeamJoinEvent:
		user, err := ConvertUserToEntity(ev.User)
		if nil != err {
			fmt.Errorf("err: ", err)
		}
		err = c.HandleTeamJoin(user)
		if nil != err {
			fmt.Errorf("err: ", err)
		}
	case *slack.UserChangeEvent:
		user, err := ConvertUserToEntity(ev.User)
		if nil != err {
			fmt.Errorf("err: ", err)
		}
		err = c.HandleUserChange(user)
		if nil != err {
			fmt.Errorf("err: ", err)
		}
	}
}
