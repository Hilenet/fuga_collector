package entity_test

import (
	"collector/domain/entity"
	"testing"
)

// 正常系
func TestNewChannel(t *testing.T) {
	slackID := "C6RNGD55L"
	name := "test"
	creator := new(entity.User)
	topic := entity.Topic{}
	purpose := entity.Purpose{}

	t.Run("generate regular channel", func(t *testing.T) {
		c, err := entity.NewChannel(slackID, name, creator, topic, purpose)

		if err != nil {
			t.Error("unexpected error:", err)
		}

		if c.SlackID != slackID {
			t.Error("channel.SlackID not exact")
		}
		if c.Name != name {
			t.Error("channel.Name not exact")
		}
		if c.Creator != creator {
			t.Error("channel.Creator not exact")
		}
		if c.Topic != topic {
			t.Error("channel.Topic not exact")
		}
		if c.Purpose != purpose {
			t.Error("channel.Purpose not exact")
		}
		if c.Archived != false {
			t.Error("channel.Archived not exact")
		}
		if c.Deleted != false {
			t.Error("channel.Deleted not exact")
		}
	})

	// t.Run("generate irregular channel", func(t *testing.T) {
	// })
}
