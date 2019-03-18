package entity_test

import (
	"collector/domain/entity"
	"testing"
	"time"
)

// 正常系
func TestNewFile(t *testing.T) {
	slackID := "FGTHGP0C8"
	filetype := "jpg"
	user := new(entity.User)
	// Unix time
	created := time.Now()
	url := "https://ritscc.slack.com/files/U5HRXRF6W/FFEJV4C9J/ios__________.jpg"

	t.Run("generate regular file", func(t *testing.T) {
		f, err := entity.NewFile(slackID, filetype, user, created, url)

		if err != nil {
			t.Error("unexpected error:", err)
		}

		if f.SlackID != slackID {
			t.Error("file.SlackID not exact")
		}
		if f.Filetype != filetype {
			t.Error("file.filetype not exact")
		}
		if f.User != user {
			t.Error("file.User not exact")
		}
		if f.Created != created {
			t.Error("file.created not exact")
		}
		if f.Url != url {
			t.Error("file.Url not exact")
		}
	})

	// t.Run("generate irregular file", func(t *testing.T) {
	// })
}
