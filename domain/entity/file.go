package entity

import (
	"collector/domain/service"
	"time"

	"github.com/pkg/errors"
)

type File struct {
	SlackID  string
	Filetype string
	User     *User
	Created  time.Time
	Url      string
}

// create entity
func NewFile(slackID string, filetype string, user *User, created time.Time, url string) (File, error) {
	if !isValidFileSlackID(slackID) || !isValidFiletype(filetype) || !isValidFileUrl(url) {
		return File{}, errors.New("[entity.#NewFile] validation err")
	}

	file := File{
		SlackID:  slackID,
		Filetype: filetype,
		User:     user,
		Created:  created,
		Url:      url,
	}

	return file, nil
}

/* === validator === */

// slackID
func isValidFileSlackID(slackID string) bool {
	if slackID[0] != 'F' {
		return false
	}
	if !service.IsValidSlackID(slackID) {
		return false
	}
	return true
}

// full listを持つほどでもねえやな
func isValidFiletype(ft string) bool {
	if len(ft) > 15 {
		return false
	}
	return true
}

// url
func isValidFileUrl(url string) bool {
	return service.IsValidUrl(url)
}
