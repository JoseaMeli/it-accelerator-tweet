package service

import (
	"github.com/IT-Accelerator-Tweet/src/domain"
	"testing"
)

func TestCanWriteATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "Async tweet")
	tweet2 := domain.NewTextTweet("grupoesfera", "Async tweet2")

	memoryTweetWriter := NewMemoryTweetWriter() //esto esta hecho para que a futuro si usamos otro medio de almacenamiento el cambio se realice solo aca, y no afuera. O sea seria una interfaz
	tweetWriter := NewChannelTweetWriter(memoryTweetWriter) // aca recibiria esa interfaz que sabria actuar dependiendo de lo que sea.

	//Channels
	tweetsToWrite := make(chan domain.Tweet)
	quit := make(chan bool)

	//Iniciar la rutina
	go tweetWriter.WriteTweet(tweetsToWrite, quit)

	// Operation
	tweetsToWrite <- tweet
	tweetsToWrite <- tweet2
	close(tweetsToWrite)

	//ESTO ACTUA COMO SEMAFORO
	<-quit

	// Validation
	if memoryTweetWriter.Tweets[0] != tweet {
		t.Errorf("A tweet in the writer was expected")
	}

	if memoryTweetWriter.Tweets[1] != tweet2 {
		t.Errorf("A tweet in the writer was expected")
	}
}
