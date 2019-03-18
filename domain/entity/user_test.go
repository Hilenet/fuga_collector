package entity_test

import (
	"collector/domain/entity"
	"testing"
)

// 正常系
func TestNewUser(t *testing.T) {
	slackID := "U02GSH6R8"
	name := "patio"
	realName := "Patio Nakamura"
	email := "hoge@gmail.com"
	iconUrl := "https://avatars.slack-edge.com//2014-08-22//2570600096_original.jpg"

	t.Run("generate regular User", func(t *testing.T) {
		u, err := entity.NewUser(slackID, name, realName, email, iconUrl)

		if err != nil {
			t.Error("unexpected error:", err)
		}

		if u.SlackID != slackID {
			t.Error("user.SlackID not exact")
		}
		if u.Name != name {
			t.Error("user.Name not exact")
		}
		if u.RealName != realName {
			t.Error("user.RealName not exact")
		}
		if u.Email != email {
			t.Error("user.Email not exact")
		}
		if u.IconUrl != iconUrl {
			t.Error("user.IconUrl not exact")
		}
		if u.Deleted != false {
			t.Error("user.Deleted not exact")
		}
	})

	// t.Run("generate irregular User", func(t *testing.T) {
	// })
}
