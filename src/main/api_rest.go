package main

import (
	"github.com/IT-Accelerator-Tweet/src/domain"
	"github.com/IT-Accelerator-Tweet/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var tm *service.TweetManager

type User struct {
	Username string `json: user`
	Password string `json: text`
}

var router *gin.Engine

func main() {

	tm = service.NewTweetManager()
	tweet1 := domain.NewTextTweet("pepe", "asdasd")
	tweet2 := domain.NewTextTweet("maw", "123345456")

	tm.PublishTweet(tweet1)
	tm.PublishTweet(tweet2)

	router = gin.Default()
	initializeRoutes()
	router.Run()




}

func initializeRoutes() {
	router.POST("/PublishTweet", PublishTweet)
	router.GET("/searchTweetById", GetTweetById)
}


func GetTweetById(c *gin.Context) {
	parametro, _ := c.GetQuery("id")
	id , _  := strconv.Atoi(parametro)

	c.String(http.StatusOK, tm.GetTweetById(id).PrintableTweet())
}

func PublishTweet(c *gin.Context) {
	var u domain.TextTweet
	c.BindJSON(&u)
	c.JSON(http.StatusOK, gin.H{
		"user": u.User,
		"pass": u.Text,
	})
}






















//
//
//
//type NewTweet struct {
//	User string `json: user`
//	Text string `json: text`
//}
//
//var router *gin.Engine
//var  tweetManager *service.TweetManager
//
//func main() {
//	tweet := domain.NewTextTweet("kei", "get1")
//	tweet2 := domain.NewTextTweet("jose", "get2")
//
//	tweetManager = service.NewTweetManager()
//
//	tweetManager.PublishTweet(tweet)
//	tweetManager.PublishTweet(tweet2)
//
//
//	router = gin.Default()
//	initializeRoutes()
//	router.Run()
//}
//
//func initializeRoutes() {
//	router.POST("/createSecretTweet", createSecretTweet)
//	router.GET("/searchTweet", getTweetById)
//	router.GET("/createTweet", createTweet)
//}
//
//func getTweetById(c *gin.Context) {
//	parametro,_ := c.GetQuery("id")
//
//	i, _ := strconv.Atoi(parametro)
//
//	c.String(http.StatusOK, tweetManager.GetTweetById(i).PrintableTweet())
//}
//
//http://localhost:8080/createTweet?user=pepe&text=lalala
//
//func createTweet(c *gin.Context) {
//	user := c.Request.URL.Query().Get("user")
//	text := c.Request.URL.Query().Get("text")
//
//	tweet := domain.NewTextTweet(user,text)
//
//	tweetManager.PublishTweet(tweet)
//
//	c.String(http.StatusOK, "Tweet >> User: "+user+ "-- Texto: "+text)
//}
//
//func createSecretTweet(c *gin.Context) {
//	var t NewTweet
//	c.BindJSON(&t)
//	c.JSON(http.StatusOK, gin.H{
//		"user": t.User,
//		"text": t.Text,
//	})
//
//	tweet := domain.NewTextTweet(t.User, t.Text)
//
//	tweetManager.PublishTweet(tweet)
//
//	c.String(http.StatusOK, "Tweet >> User: "+t.User+ "-- Texto: "+t.Text)
//}
//
