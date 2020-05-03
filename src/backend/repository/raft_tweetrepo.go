package repository

import (
	"context"
	"encoding/json"
	"time"
)

func SaveTweet(tweetUser string,tweetContent string,ctx context.Context)(error){
	resultChan := make(chan bool)
	deleteChan := make(chan error)
	go SaveTweetDB(tweetUser,tweetContent,resultChan,deleteChan,ctx)
	select {
	case <-resultChan:
		return nil
	case err := <-deleteChan:
		return err
	}
}

func SaveTweetDB(tweetUser string,tweetContent string,resultChan chan bool, deleteChan chan error, ctx context.Context){
	tweetContent += "*"+time.Now().Format("2006-01-02 15:04:05")
	var tweetList []string
	res,geterr := Get("%"+tweetUser,ctx)
	if geterr != nil{
		deleteChan <- geterr
	}else{
		if string(res) == ""{
			tweetList = make([]string,0)
		}else {
			err:= json.Unmarshal(res,&tweetList)
			if err != nil{
				deleteChan <- err
			}
		}
	}
	tweetList = append(tweetList,tweetContent)
	value,err := json.Marshal(tweetList)
	if err != nil{
		deleteChan <- err
	}
	puterr := Put("%"+tweetUser,string(value),ctx)
	if puterr != nil{
		deleteChan <- puterr
	}else {
		resultChan <- true
	}
}

func GetTweetList(followUsername string,ctx context.Context)([]string,error) {
	resultChan := make(chan []string)
	deleteChan := make(chan error)
	dummyList := make([]string,0)
	go GetTweetListDB(followUsername,resultChan,deleteChan,ctx)
	select {
	case res := <-resultChan:
		return res,nil
	case err := <-deleteChan:
		return dummyList, err
	}
}

func GetTweetListDB(followUsername string, resultChan chan []string,deleteChan chan error, ctx context.Context){
	var tweetList []string
	res,geterr := Get("%"+followUsername,ctx)
	if geterr != nil{
		deleteChan <- geterr
	}else{
		if string(res) == ""{
			tweetList = make([]string,0)
			resultChan <- tweetList
		}else{
			err:= json.Unmarshal(res,&tweetList)
			if err != nil{
				deleteChan <- err
			}else {
				resultChan <- tweetList
			}
		}

	}
}
