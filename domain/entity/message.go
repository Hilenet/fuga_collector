package entity

import (
	"collector/domain/service"

	"github.com/pkg/errors"
)

// WIP
type Message struct {
	// Unixtime.\d{6}
	Ts      string
	Channel *Channel
	User    *User
	Text    string
	Edited  bool
	Deleted bool
	// attachmentsなんざ知らねー！
	// json, nullableなslice
	Attachments []string
	// nullable
	SubAttribute SubAttribute
}

// nullable,
type SubAttribute struct {
	Subtype string
}

func NewMessage(ts string, channel *Channel, user *User, text string, attachments []string) (Message, error) {
	if !isValidMessageTimestamp(ts) || !isValidMessageText(text) {
		return Message{}, errors.New("[entity.#NewMessage] validation err")
	}
	message := Message{
		Ts:          ts,
		Channel:     channel,
		User:        user,
		Text:        text,
		Attachments: attachments,
		Edited:      false,
		Deleted:     false,
	}

	return message, nil
}

/* === validator === */

// ts(as string)
func isValidMessageTimestamp(str string) bool {
	return service.IsValidSlackTimestamp(str)
}

// text
func isValidMessageText(text string) bool {
	if len(text) > 65535 {
		return false
	}
	return true
}
