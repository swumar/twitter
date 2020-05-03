package repository

import (
	"backend/model"
	"context"
	"encoding/json"
	//"fmt"
)

func ReturnUser(username string, ctx context.Context)(model.User, bool, error){
	resultChan := make(chan model.User)
	errChan := make(chan bool)
	deleteChan := make(chan error)
	dummy := new(model.User)
	dummyUser := *dummy
	go ReturnUserDB(username,resultChan,errChan,deleteChan,ctx)

	select {
	case res := <-resultChan :
		return res, <-errChan, nil
	case err :=<-deleteChan:
		return dummyUser,false,err
	}

}

func ReturnUserDB(username string, resultChan chan model.User, errChan chan bool,deleteChan chan error, ctx context.Context)  {
	var user model.User
	res,geterr := Get("#"+username,ctx)
	if geterr != nil{
		deleteChan <- geterr
	}else{
		err:= json.Unmarshal(res,&user)
		if err != nil{
			deleteChan <- geterr
		}else {
			if user.Username != ""{
				resultChan <- user
				errChan <- true
			}else {
				resultChan <- user
				errChan <- false
			}
		}
	}
}

func SaveUser(user model.User, ctx context.Context)(error){
	resultChan := make(chan bool)
	deleteChan := make(chan error)
	go SaveUserDB(user,resultChan,deleteChan,ctx)

	select {
	case <-resultChan:
		return nil
	case err := <-deleteChan:
		return err
	}

}

func SaveUserDB(user model.User,resultChan chan bool,deleteChan chan error,ctx context.Context)  {
	value,err := json.Marshal(user)
	if err != nil{
		deleteChan <- err
	}
	puterr := Put("#"+user.Username,string(value),ctx)
	if puterr != nil{
		deleteChan <- puterr
	}else {
		resultChan <- true
	}
}

func GetUsers(ctx context.Context)(string,error){
	resultChan := make(chan string)
	deleteChan := make(chan error)
	go GetUsersDB(resultChan,deleteChan,ctx)
	select {
	case res := <-resultChan:
		return res,nil
	case err := <-deleteChan:
		return "",err
	}
}

func GetUsersDB(resultChan chan string,deleteChan chan error, ctx context.Context)  {
	res,geterr := GetAll(ctx)
	if geterr != nil{
		deleteChan <- geterr
	}else{
		resultChan <- res
	}
}
