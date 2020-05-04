package repository

import (
	"context"
	"encoding/json"
	"strconv"
	"sync"
	"testing"
	"time"
)

//Adding 10 users with a tweet in TestSaveTweet which will be used as mockup data for other test cases.
//So please run the test cases in the same order if running it individually.

func TestSaveTweet(t *testing.T) {
	Delete(context.Background())
	wg := sync.WaitGroup{}
	for i := 0 ; i < 10 ; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			tweetUser := "user" + strconv.Itoa(v)
			tweetContent := "tweet" + strconv.Itoa(v)
			err := SaveTweet(tweetUser,tweetContent,context.Background())
			if err != nil {
				t.Log("Problem in saving tweet of user", v)
			}
		}(i)
	}
	wg.Wait()

	for i := 0 ; i < 10 ; i++ {
		var tweetList []string
		tweetUser := "user" + strconv.Itoa(i)
		res,_ := Get("%"+tweetUser,context.Background())
		_ = json.Unmarshal(res,&tweetList)
		if len(tweetList) != 1{
			t.Error("Error while saving tweet of"+tweetUser)
		}else {
			t.Log(tweetList)
		}
	}
	t.Log("Test SaveTweet successful")
}

func TestGetTweetList(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		tweetUser := "user0"
		tweetList, err := GetTweetList(tweetUser,context.Background())
		if err != nil {
			t.Error("Problem in getting tweets of user")
		}else {
			t.Log(tweetList)
			t.Log("Test GetTweetList successful")
		}
		}()
	wg.Wait()
}

func TestGetTweetListContext(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		tweetUser := "user0"
		ctx , cancel := context.WithTimeout(context.Background(),time.Duration(100)*time.Second)
		cancel()
		tweetList, err := GetTweetList(tweetUser,ctx)
		if err == nil {
			t.Log(tweetList)
			t.Error("Test GetTweetList unsuccessful")
		}else {
			t.Log("Test GetTweetList successful")
		}
	}()
	wg.Wait()
}

func TestSaveTweetContext(t *testing.T) {
	Delete(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx , cancel := context.WithTimeout(context.Background(),time.Duration(100)*time.Second)
		cancel()
		tweetUser := "user0"
		tweetContent := "tweet0"
		err := SaveTweet(tweetUser,tweetContent,ctx)
		if err == nil {
			t.Error("Test SaveTweetContext unsuccessful")
		}else {
			t.Log("Test SaveTweetContext successful")
		}
	}()
	wg.Wait()
}