package service

import (
	"net/url"
	"regexp"
)

/* === validator helper === */

// slackID
var slackID_pattern = `^[\da-zA-Z]{9}$`
var slackID_reg = regexp.MustCompile(slackID_pattern)

func IsValidSlackID(slackID string) bool {
	return slackID_reg.MatchString(slackID)
}

// slack timestamp
var timestamp_pattern = `^\d{10}.\d{6}$`
var timestamp_reg = regexp.MustCompile(timestamp_pattern)

func IsValidSlackTimestamp(ts string) bool {
	return timestamp_reg.MatchString(ts)
}

// email
var email_pattern = `^(?i:[^ @"<>]+|".*")@(?i:[a-z1-9.])+.(?i:[a-z])+$`
var email_reg = regexp.MustCompile(email_pattern)

func IsValidEmail(email string) bool {
	return email_reg.MatchString(email)
}

// url，net/urlのparse信用できねえな
func IsValidUrl(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}
