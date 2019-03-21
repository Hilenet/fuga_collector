package handler

import (
	"collector/domain/entity"
	"os"

	"github.com/nlopes/slack"
)

type SlackHandler struct {
	Client *slack.Client
	RTM    *slack.RTM
}

func NewSlackHandler(token string) (SlackHandler, error) {
	client := slack.New(token)
	// この依存は必ず殺せ
	if os.Getenv("PROD_MODE") == "development" {
		client.SetDebug(true)
	}
	_, err := client.AuthTest()
	if err != nil {
		return SlackHandler{}, err
	}

	rtm := client.NewRTM()
	// connectと初期化処理?
	go rtm.ManageConnection()

	slackHandler := SlackHandler{
		Client: client,
		RTM:    rtm,
	}

	return slackHandler, nil
}

func ConvertUserToEntity(su slack.User) (entity.User, error) {
	user, err := entity.NewUser(su.ID, su.Name, su.RealName, su.Profile.Email, su.Profile.ImageOriginal)

	return user, err
}
