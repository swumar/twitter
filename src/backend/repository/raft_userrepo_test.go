package repository

import (
	"backend/model"
	"context"
	"strconv"
	"strings"
	"sync"
	"testing"
	//"time"
)

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
		t.Log("Test SaveUser succesful")
	}else {
		t.Error("Test case unsucessful")
	}

}
