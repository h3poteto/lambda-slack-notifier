package main

import (
	"time"
)

type Event struct {
	AccountID string
	UserName  string
	EventType string
	Result    *ResponseElements
	IPAddress string
	Time      time.Time
}

func NewEvent(accountID, userName, eventType string, result *ResponseElements, ipAddress string, time time.Time) *Event {
	return &Event{
		accountID,
		userName,
		eventType,
		result,
		ipAddress,
		time,
	}
}

// PostSlack post the event to a slack channel
func (e *Event) PostSlack(url, channel string) error {
	s := NewSlack(url)
	return s.postEvent(e, channel)
}

type Detail struct {
	SourceIPAddress  string            `json:"sourceIPAddress"`
	EventType        string            `json:"eventType"`
	UserIdentity     *UserIdentity     `json:"userIdentity"`
	ResponseElements *ResponseElements `json:"responseElements"`
}

type UserIdentity struct {
	Type     string `json:"type"`
	UserName string `json:"userName"`
}

type ResponseElements struct {
	ConsoleLogin string `json:"ConsoleLogin"`
	CheckMfa     string `json:"CheckMfa"`
}
