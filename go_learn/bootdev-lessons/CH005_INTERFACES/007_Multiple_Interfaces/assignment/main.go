package main

import (
	"fmt"
)

func (e email) cost() int {
	if !e.isSubscribed {
		return len(e.body) * 5
	}
	return len(e.body) * 2
}

func (e email) format() string {
	subscribedText := "Subscribed"
	if !e.isSubscribed {
		subscribedText = "Not Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, subscribedText)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
