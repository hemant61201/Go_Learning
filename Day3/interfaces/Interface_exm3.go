package main

import (
	"fmt"
)

func (e email) cost() int {
	if e.isSubscribed {
		return 2 * len(e.body)
	}

	return 5 * len(e.body)
}

func (e email) format() string {
	if e.isSubscribed {
		return fmt.Sprintf("'%s' | Subscribed", e.body)
	}
	return fmt.Sprintf("'%s' | Not Subscribed", e.body)
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
