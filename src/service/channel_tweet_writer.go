package service

import "github.com/IT-Accelerator-Tweet/src/domain"

type ChannelTweetWriter struct {
	MemoryTweetWritter  *MemoryTweetWriter
}


func NewChannelTweetWriter(tw *MemoryTweetWriter) *ChannelTweetWriter{
	ctw := ChannelTweetWriter{}
	ctw.MemoryTweetWritter = tw

	return &ctw
}

func (tw *ChannelTweetWriter) WriteTweet(channelTweet chan domain.Tweet, quit chan bool) {

	for tweet := range channelTweet {
		tw.MemoryTweetWritter.Tweets = append(tw.MemoryTweetWritter.Tweets, tweet)
	}
	quit<-true
}