package domain

import (
	"fmt"
	"time"
)

type TextTweet struct {
	Id   int
	User string
	Text string
	Date *time.Time
}
var ContadorDeTweets int

func NewTextTweet(user string, text string) *TextTweet {
	t := time.Now()
	tweet := TextTweet{ContadorDeTweets,user,text, &t}

	var puntero *TextTweet
	puntero = &tweet

	ContadorDeTweets ++

	return puntero
}



func (t *TextTweet) String() string {
	return fmt.Sprintf("@%s: %s", t.User, t.Text)
}

func (t *TextTweet) GetUser() string {
	return t.User
}

func (t *TextTweet) GetText() string {
	return t.Text
}

func (t *TextTweet) GetId() int {
	return t.Id
}

func (t *TextTweet) GetDate() *time.Time {
	return t.Date
}


func (t *TextTweet) PrintableTweet() string{
	return fmt.Sprintf("@%s: %s", t.User, t.Text)
}