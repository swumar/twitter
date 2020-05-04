package repository

import (
	"backend/model"
	"context"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

//Adding 10 users in TestSaveUser which will be used as mockup data for other test cases. So please run the test cases
//in the same order if running it individually.

func TestSaveUser(t *testing.T) {
	Delete(context.Background())
	wg := sync.WaitGroup{}
	for i:=0 ; i < 10 ; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()

			user := model.User{
				Username:  "test1_"+strconv.Itoa(v)+"@gmail.com",
				Password:  "1234",
				FirstName: "us"+strconv.Itoa(v),
				LastName:  "er"+strconv.Itoa(v),
			}
			err := SaveUser(user,context.Background())
			if err != nil {
				t.Error("Problem in adding user")
			}
		}(i)
	}
	wg.Wait()
	res,_ := GetAll(context.Background())
	res = res[:len(res)-1]
	t.Log(res)
	resSlice := strings.Split(res,",")
	if len(resSlice) == 10{
		t.Log("Test SaveUser successful")
	}else {
		t.Error("Test SaveUser unsuccessful")
	}

}

func TestReturnUser(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		username := "test1_0@gmail.com" //added in TestSaveUser test case
		user, _, err := ReturnUser(username,context.Background())
		if err != nil {
			t.Error("Test ReturnUser unsuccessful")
		}else {
			t.Log(user.Username)
			t.Log("Test ReturnUser successful")
		}
	}()
	wg.Wait()
}

func TestReturnUserContext(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		username := "test1_0@gmail.com" //added in TestSaveUser test case
		ctx , cancel := context.WithTimeout(context.Background(),time.Duration(100)*time.Second)
		cancel()
		user, _, err := ReturnUser(username,ctx)
		if err != nil {
			t.Log("Test ReturnUserContext successful")
		}else {
			t.Log(user.Username)
			t.Error("Test ReturnUserContext unsuccessful")
		}
	}()
	wg.Wait()
}

func TestGetUsers(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		users, err := GetUsers(context.Background())
		if err != nil {
			t.Error("Test GetUsers unsuccessful")
		}else {
			t.Log(users)
			t.Log("Test GetUsers successful")
		}
	}()
	wg.Wait()
}

func TestGetUsersContext(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx , cancel := context.WithTimeout(context.Background(),time.Duration(100)*time.Second)
		cancel()
		users, err := GetUsers(ctx)
		if err != nil {
			t.Log("Test GetUsersContext successful")
		}else {
			t.Log(users)
			t.Error("Test GetUsersContext unsuccessful")
		}
	}()
	wg.Wait()
}

func TestSaveUserContext(t *testing.T) {
	Delete(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		user := model.User{
			Username:  "test2_1@gmail.com",
			Password:  "1234",
			FirstName: "us1",
			LastName:  "er1",
		}
		ctx , cancel := context.WithTimeout(context.Background(),time.Duration(100)*time.Second)
		cancel()
		err := SaveUser(user,ctx)
		if err == nil {
			t.Error("User added successfully")
		}
	}()
	wg.Wait()
	res,_ := GetAll(context.Background())
	if len(res) == 0{
		t.Log("Test SaveUserContext successful")
	}else {
		t.Log(res)
		t.Error("Test SaveUserContext unsuccessful")
	}

}