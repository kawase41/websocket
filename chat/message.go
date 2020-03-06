package main

import (
	"time"
)

// message represents one message
type message struct {
	Name string
	Message string
	When time.Time
	AvatarURL string
}