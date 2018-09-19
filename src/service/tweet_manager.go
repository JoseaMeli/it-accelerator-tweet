package service

import (
	"github.com/IT-Accelerator-Tweet/src/domain"
	"github.com/pkg/errors"
)

type TweetManager struct {
	Tweet domain.Tweet
	ContenedorDeTweets []domain.Tweet
	Um *UsuarioManager
}

//-----------FUNCION--------------//

func NewTweetManager() *TweetManager{
	tm := TweetManager{}
	tm.ContenedorDeTweets = make([]domain.Tweet,0)
	tm.Um = NewUsuarioManager()

	//tm.Um.AgregarUsuario(&domain.Usuario{"kei", "kei@meli.com", "kei", "123"})
	//tm.Um.AgregarUsuario(&domain.Usuario{"jose", "jose@meli.com", "jose", "1234"})
	return &tm
}

//----------METODOS--------------//

func (tm *TweetManager)PublishTweet(tweet domain.Tweet) (int, error) {

	//usuarioEncontrado := tm.Um.GetUsuarioByNombre(tweet.GetUser())

	//if usuarioEncontrado != nil {
		if tweet.GetUser() == "" {
			return 0, errors.New("user is required")
		} else if tweet.GetText() == "" {
			return 0, errors.New("text is required")
		} else if len(tweet.GetText()) > 140 {
			return 0, errors.New("text exceeds 140 characters")
		} else {
			tm.Tweet = tweet
			tm.ContenedorDeTweets = append(tm.ContenedorDeTweets, tweet)
			return tweet.GetId(), nil
		}
	//} else {
	//	return 0, errors.New("No existe el Usuario")
	//}
}

func (tm *TweetManager)GetTweet() domain.Tweet {
	return tm.Tweet
}

func (tm *TweetManager)GetTweets() []domain.Tweet {
	return tm.ContenedorDeTweets
}

func (tm *TweetManager)GetTweetById(id int)  domain.Tweet{
	for _,tweet := range tm.ContenedorDeTweets {
		if tweet.GetId() == id {
			return tweet
		}
	}
	return nil
}

func (tm *TweetManager)CountTweetsByUser(user string) int {
	var i int
	for _,tweet := range tm.ContenedorDeTweets {
		if tweet.GetUser() == user {
			i ++
		}
	}
	return i
}

func (tm *TweetManager)GetTweetsByUser(user string) []domain.Tweet {
	tweets := make([]domain.Tweet,0)
	for _,tweet := range tm.ContenedorDeTweets {
		if tweet.GetUser() == user {
			tweets = append(tweets, tweet)
		}
	}
	return tweets
}