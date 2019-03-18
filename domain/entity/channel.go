package entity

import (
	"collector/domain/service"

	"github.com/pkg/errors"
)

type Channel struct {
	SlackID  string
	Name     string
	Creator  *User
	Topic    Topic
	Purpose  Purpose
	Archived bool
	Deleted  bool
}

// デフォ値はgoのゼロ値でおk
type Topic struct {
	Value   string
	Creator *User
	LastSet int
}

// デフォ値はgoのゼロ値でおk
type Purpose struct {
	Value   string
	Creator *User
	LastSet int
}

// create entity from raw values
func NewChannel(slackID string, name string, creator *User, topic Topic, purpose Purpose) (Channel, error) {
	// validation
	if !isValidChannelSlackID(slackID) || !isValidChannelName(name) {
		return Channel{}, errors.New("[entity.#NewChannel] validation err")
	}

	channel := Channel{
		SlackID:  slackID,
		Name:     name,
		Creator:  creator,
		Topic:    topic,
		Purpose:  purpose,
		Archived: false,
		Deleted:  false,
	}

	return channel, nil
}

/* === validator === */

// slackID
func isValidChannelSlackID(slackID string) bool {
	if slackID[0] != 'C' {
		return false
	}
	if !service.IsValidSlackID(slackID) {
		return false
	}
	return true
}

// name
func isValidChannelName(name string) bool {
	if len(name) > 255 {
		return false
	}
	return true
}
