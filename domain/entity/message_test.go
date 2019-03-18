package entity_test

import (
	"collector/domain/entity"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func readJson(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()

	buf := make([]byte, 1024)
	builder := new(strings.Builder)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}

		fmt.Fprint(builder, buf)
	}
	return builder.String()
}

// 正常系
func TestNewmessage(t *testing.T) {
	ts := "1550916487.000000"
	channel := new(entity.Channel)
	user := new(entity.User)
	text := "hoge"
	attachments := []string{
		readJson("testdata/message_test01.json"),
	}

	t.Run("generate regular message", func(t *testing.T) {
		// TOOD: newしたスライスの参照先arrayはどのスコープで生存するのか
		m, err := entity.NewMessage(ts, channel, user, text, attachments)

		if err != nil {
			t.Error("unexpected error:", err)
		}

		if m.Ts != ts {
			t.Error("message.Ts not exact")
		}
		if m.Channel != channel {
			t.Error("message.Channel not exact")
		}
		if m.User != user {
			t.Error("message.User not exact")
		}
		if m.Text != text {
			t.Error("message.Text not exact")
		}
		// ここもうちょいなんとかなりそう
		if !reflect.DeepEqual(m.Attachments, attachments) {
			t.Error("message.Attachments not exact")
		}
		if m.Edited != false {
			t.Error("message.Edited not exact")
		}
		if m.Deleted != false {
			t.Error("message.Deleted not exact")
		}
	})

	// t.Run("generate irregular message", func(t *testing.T) {
	// })
}
