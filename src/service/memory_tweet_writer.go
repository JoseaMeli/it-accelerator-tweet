package service

import "github.com/IT-Accelerator-Tweet/src/domain"

type MemoryTweetWriter struct {
	Tweets []domain.Tweet
}

func NewMemoryTweetWriter() *MemoryTweetWriter{
	mtw := MemoryTweetWriter{}
	mtw.Tweets = make([]domain.Tweet,0)

	return &mtw
}