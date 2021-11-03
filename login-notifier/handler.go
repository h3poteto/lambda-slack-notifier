package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	SlackURL     string `envconfig:"SLACK_URL" required:"true"`
	SlackChannel string `envconfig:"SLACK_CHANNEL" required:"true"`
}

func handler(ctx context.Context, event events.CloudWatchEvent) {
	var env EnvConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Fatal(err)
		return
	}

	// for debug logging
	var d interface{}
	json.Unmarshal(event.Detail, &d)
	log.Printf("[DEBUG] detail: %+v\n", d)

	var detail *Detail
	err := json.Unmarshal(event.Detail, &detail)
	if err != nil {
		log.Fatal(err)
		return
	}
	e := NewEvent(
		event.AccountID,
		detail.UserIdentity.UserName,
		detail.EventType,
		detail.ResponseElements,
		detail.SourceIPAddress,
		event.Time,
	)

	if err := e.PostSlack(env.SlackURL, env.SlackChannel); err != nil {
		log.Fatal(err)
	}
	return
}
