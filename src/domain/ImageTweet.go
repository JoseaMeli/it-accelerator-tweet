package domain

import (
	"fmt"
	"time"
)

type ImageTweet struct {
	TextTweet
	Url string
}

func NewImageTweet(user string, text string, url string) *ImageTweet {
	t := time.Now()
	tt := TextTweet{ContadorDeTweets,user,text, &t}
	tweet := ImageTweet{tt, url}

	var puntero *ImageTweet
	puntero = &tweet

	ContadorDeTweets ++

	return puntero
}

func (t *ImageTweet) PrintableTweet() string{
	return fmt.Sprintf("@%s: %s %s", t.User, t.Text,t.Url)
}

func (t *ImageTweet) String() string {
	return fmt.Sprintf("@%s: %s %s", t.User, t.Text,t.Url)
}
