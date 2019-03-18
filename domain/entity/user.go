package entity

import (
	"collector/domain/service"

	"github.com/pkg/errors"
)

type User struct {
	SlackID  string
	Name     string
	RealName string
	Email    string
	IconUrl  string
	Deleted  bool
}

// create entity from raw values
func NewUser(slackID string, name string, realName string, email string, iconUrl string) (User, error) {
	// validation
	if !isValidUserSlackID(slackID) || !isValidUserName(name) || !isValidUserName(realName) || !isValidUserEmail(email) || !isValidUserIconUrl(iconUrl) {
		return User{}, errors.New("[entity.#NewUser] validation err")
	}

	user := User{
		SlackID:  slackID,
		Name:     name,
		RealName: realName,
		Email:    email,
		IconUrl:  iconUrl,
		Deleted:  false,
	}

	return user, nil
}

/* === validator === */

// slackID
func isValidUserSlackID(slackID string) bool {
	if slackID[0] != 'U' {
		return false
	}
	if !service.IsValidSlackID(slackID) {
		return false
	}
	return true
}

// name, realname
func isValidUserName(name string) bool {
	if len(name) > 255 {
		return false
	}
	return true
}

// email
func isValidUserEmail(email string) bool {
	return service.IsValidEmail(email)
}

// iconUrl
func isValidUserIconUrl(iconUrl string) bool {
	return service.IsValidUrl(iconUrl)
}
