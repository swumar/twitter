package repository

import (
	"backend/model"
	"context"
	"time"
)

var Tweets = make(map[string][]string)

func SaveTweet(tweetUser string,tweetContent string,ctx context.Context)(error){
	resultChan := make(chan bool)
	deleteChan := make(chan bool)
	go SaveTweetDB(tweetUser,tweetContent,resultChan,deleteChan,ctx)
	select {
	case <-resultChan:
		return nil
	case <-deleteChan:
		return ctx.Err()
	}
}

func SaveTweetDB(tweetUser string,tweetContent string,resultChan chan bool, deleteChan chan bool, ctx context.Context){
	tweetContent += "*"+time.Now().Format("2006-01-02 15:04:05")
	model.TweetsMux.Lock()
	if Tweets[tweetUser] == nil{
		Tweets[tweetUser] = make([]string,0)
	}
	Tweets[tweetUser] = append(Tweets[tweetUser],tweetContent)

	select {
	case <-ctx.Done():
		model.TweetsMux.Unlock()
		channel := make(chan bool)
		go DeleteTweetDB(tweetUser,channel)
		<-channel
		deleteChan <- true
	default:
		model.TweetsMux.Unlock()
		resultChan <- true
	}
}

func DeleteTweetDB(tweetUser string,resultChan chan bool) {
	model.TweetsMux.Lock()
	Tweets[tweetUser] = Tweets[tweetUser][:len(Tweets[tweetUser])-1]
	model.TweetsMux.Unlock()
	resultChan <- true
}

func GetTweetList(followUsername string,ctx context.Context)([]string,error) {
	resultChan := make(chan []string)
	deleteChan := make(chan bool)
	dummyList := make([]string,0)
	go GetTweetListDB(followUsername,resultChan,deleteChan,ctx)
	select {
	case res := <-resultChan:
		return res,nil
	case <-deleteChan:
		return dummyList, ctx.Err()
	}
}

func GetTweetListDB(followUsername string, resultChan chan []string,deleteChan chan bool, ctx context.Context){
	model.TweetsMux.Lock()
	tweetList := Tweets[followUsername]
	select {
	case <-ctx.Done():
		model.TweetsMux.Unlock()
		deleteChan <- true
	default:
		model.TweetsMux.Unlock()
		resultChan <- tweetList
	}
}
