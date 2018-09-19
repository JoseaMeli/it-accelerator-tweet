package domain

import "time"

type Tweet interface {
	PrintableTweet() string
	GetUser() string
	GetText() string
	GetId()	int
	GetDate() *time.Time
}
