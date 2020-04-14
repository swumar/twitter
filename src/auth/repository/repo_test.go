package repository

import (
	"container/list"
	"context"
	"strconv"
	"sync"
	"testing"
	authStorage "auth/storage"
	authmodel "auth/model"
	"time"
)

//Test for SaveUser is same as test for SaveUserRegister without context cancelling

func TestSaveUserRegister(t *testing.T) {

	wg := sync.WaitGroup{}
	for i:=0 ; i < 10 ; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()

			user := authmodel.User{
				Username:  "user"+strconv.Itoa(v)+"@gmail.com",
				Password:  "1234",
				FirstName: "us"+strconv.Itoa(v),
				LastName:  "er"+strconv.Itoa(v),
				Followers: list.New(),
			}
			err := SaveUserRegister(user,context.Background())
			if err != nil {
				t.Error("Problem in adding user")
			}
		}(i)
	}
	wg.Wait()
	if len(authStorage.Users) == 10{
		t.Log("Test SaveUserRegister succesful")
	}else{
		t.Errorf("Number of users missing %d",10-len(authStorage.Users))
	}
}

func TestSaveUserRegisterContext(t *testing.T) {

	wg := sync.WaitGroup{}
	for i:=0 ; i < 10 ; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()

			user := authmodel.User{
				Username:  "user"+strconv.Itoa(v)+"@gmail.com",
				Password:  "1234",
				FirstName: "us"+strconv.Itoa(v),
				LastName:  "er"+strconv.Itoa(v),
				Followers: list.New(),
			}
			if v % 2 == 0{
				ctx , cancel := context.WithTimeout(context.Background(),time.Duration(100)*time.Second)
				defer cancel()
				err := SaveUserRegister(user,ctx)
				if err != nil {
					t.Log("Problem in adding user:",v)
				}else {
					t.Log("User added sucessfully:",v)
				}
			}else{
				ctx , cancel := context.WithTimeout(context.Background(),time.Duration(1)*time.Millisecond)
				defer cancel()
				err := SaveUserRegister(user,ctx)
				if err != nil {
					t.Log("Problem in adding user:",v)
				}else {
					t.Log("User added sucessfully:",v)
				}
			}

		}(i)
	}
	wg.Wait()
	t.Log(authStorage.Users)
	if len(authStorage.Users) == 5{
		t.Log("Test SaveUserRegister succesful")
	}else{
		t.Errorf("Number of users missing %d",5-len(authStorage.Users))
	}
}

func TestSaveUserContext(t *testing.T) {

	wg := sync.WaitGroup{}
	for i:=0 ; i < 10 ; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()

			user := authmodel.User{
				Username:  "user"+strconv.Itoa(v)+"@gmail.com",
				Password:  "1234",
				FirstName: "us"+strconv.Itoa(v),
				LastName:  "er"+strconv.Itoa(v),
				Followers: list.New(),
			}
			bkpUser := user
			user.Token = "xxxxxxxxxxxx"
			if v % 2 == 0{
				ctx , cancel := context.WithTimeout(context.Background(),time.Duration(100)*time.Second)
				defer cancel()
				err := SaveUser(user,ctx,bkpUser)
				if err != nil {
					t.Log("Problem in modifying user:",v)
				}else {
					t.Log("User modified sucessfully:",v)
				}
			}else{
				ctx , cancel := context.WithTimeout(context.Background(),time.Duration(1)*time.Millisecond)
				defer cancel()
				err := SaveUser(user,ctx,bkpUser)
				if err != nil {
					t.Log("Problem in modifying user:",v)
				}else {
					t.Log("User modified sucessfully:",v)
				}
			}

		}(i)
	}
	wg.Wait()
	t.Log(authStorage.Users)
	count := 0
	for i:=0 ; i<10 ; i++{
		user := authStorage.Users["user"+strconv.Itoa(i)+"@gmail.com"]
		if user.Token != ""{
			count = count + 1
		}
	}
	if count == 5{
		t.Log("Test SaveUserRegister succesful")
	}else{
		t.Errorf("Error in modifying users")
	}
}