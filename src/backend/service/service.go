package main

import (
	"backend/model"
	"backend/proto"
	"backend/repository"
	"context"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (*server) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {

	user, usernameExists, ctxErr := repository.ReturnUser(request.Username,ctx)

	if ctxErr != nil{
		response := &proto.LoginResponse{Message:"Request timeout. Try again", Tokenstring: ""}
		return response, nil
	}

	if usernameExists {
		passowrdErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

		if passowrdErr != nil {
			response := &proto.LoginResponse{Message:"Invalid Password", Tokenstring: ""}
			return response, nil
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username":  user.Username,
			"firstname": user.FirstName,
			"lastname":  user.LastName,
		})

		tokenString, loginErr := token.SignedString([]byte("secret"))

		if loginErr != nil {
			response := &proto.LoginResponse{Message:"Error while generating token,Try again", Tokenstring: ""}
			return response, nil
		}

		user.Token = tokenString

		ctxErr2 := repository.SaveUser(user,ctx)

		if ctxErr2 != nil{
			response := &proto.LoginResponse{Message:"Request timeout. Try again", Tokenstring: ""}
			return response, nil
		}

		response := &proto.LoginResponse{Message:"", Tokenstring: tokenString}
		return response, nil

	}

	response := &proto.LoginResponse{Message:"Invalid Username", Tokenstring: ""}
	return response, nil

}

func (*server) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.RegisterResponse, error) {

	_, usernameExists, ctxErr := repository.ReturnUser(request.Username,ctx)
	if ctxErr != nil{
		response := &proto.RegisterResponse{Message:"Request timeout. Try again"}
		return response, ctxErr
	}
	if usernameExists {
		response := &proto.RegisterResponse{Message: "User already exists"}
		return response, nil
	}

	registerFromInput := model.User{
		Username:  request.Username,
		Password:  request.Password,
		FirstName: request.Firstname,
		LastName:  request.Lastname,
		Followers: make([]string,0),
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(registerFromInput.Password), 5)
	if err != nil {
		response := &proto.RegisterResponse{Message: "Error While Hashing Password, Try Again"}
		return response, nil
	}
	registerFromInput.Password = string(hash)

	ctxErr2 := repository.SaveUser(registerFromInput,ctx)

	if ctxErr2 != nil{
		response := &proto.RegisterResponse{Message:"Request timeout. Try again"}
		return response, ctxErr2
	}

	response := &proto.RegisterResponse{Message: "",}
	return response, nil
}

func (*server) Logout(ctx context.Context, request *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	token, tokenerr := jwt.Parse(request.Tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if !token.Valid || tokenerr != nil {
		response := &proto.LogoutResponse{Message: "Please login to continue"}
		return response, nil
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	signoutUserName := claims["username"].(string)
	signoutUser, _, ctxErr := repository.ReturnUser(signoutUserName,ctx)

	if ctxErr != nil{
		response := &proto.LogoutResponse{Message: "Request timeout. Try again"}
		return response, ctxErr
	}

	if signoutUser.Username != "" {
		signoutUser.Token = ""
		ctxErr2 := repository.SaveUser(signoutUser,ctx)
		if ctxErr2 != nil{
			response := &proto.LogoutResponse{Message: "Request timeout. Try again"}
			return response, ctxErr2
		}
		response := &proto.LogoutResponse{Message: ""}
		return response, nil
	} else {
		response := &proto.LogoutResponse{Message: "Please login to continue"}
		return response, nil
	}
}

func (*server) FollowService(ctx context.Context, request *proto.ProfileRequest) (*proto.ProfileResponse, error) {
	userPresent, _, ctxErr1 := repository.ReturnUser(request.GetReqparm1(),ctx)
	if ctxErr1 != nil{
		response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
		return response, ctxErr1
	}

	followUser, _, ctxErr2 := repository.ReturnUser(request.GetReqparm2(),ctx)

	if ctxErr2 != nil{
		response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
		return response, ctxErr2
	}
	if userPresent.Username == followUser.Username{
		response := &proto.ProfileResponse{Resparm1: "Cant follow yourself"}
		return response, nil
	}

	for i := 0 ; i < len(followUser.Followers) ; i++{
		if userPresent.Username == followUser.Followers[i]{
			response := &proto.ProfileResponse{Resparm1: "User already followed"}
			return response, nil
		}
	}

	if userPresent.Username != "" {
		followUser.Followers = append(followUser.Followers,userPresent.Username)
		ctxErr3 := repository.SaveUser(followUser,ctx)
		if ctxErr3 != nil{
			response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
			return response, ctxErr3
		}
		response := &proto.ProfileResponse{Resparm1: ""}
		return response, nil
	} else {
		response := &proto.ProfileResponse{Resparm1: "Username doesnt exist"}
		return response, nil
	}

}

func (*server) UnfollowService(ctx context.Context, request *proto.ProfileRequest) (*proto.ProfileResponse, error) {

	userPresent, _ , ctxErr1 := repository.ReturnUser(request.GetReqparm1(),ctx)

	if ctxErr1 != nil{
		response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
		return response, ctxErr1
	}

	unfollowUser, _, ctxErr2 := repository.ReturnUser(request.GetReqparm2(),ctx)

	if ctxErr2 != nil{
		response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
		return response, ctxErr2
	}

	if userPresent.Username == unfollowUser.Username{
		response := &proto.ProfileResponse{Resparm1: "Cant unfollow yourself"}
		return response, nil
	}

	if userPresent.Username == "" {
		response := &proto.ProfileResponse{Resparm1: "Username doesnt exist"}
		return response, nil	}

	for i := 0 ; i < len(unfollowUser.Followers) ; i++{
		if userPresent.Username == unfollowUser.Followers[i]{
			unfollowUser.Followers = append(unfollowUser.Followers[:i],unfollowUser.Followers[i+1:]...)
			ctxErr3 := repository.SaveUser(unfollowUser,ctx)
			if ctxErr3 != nil{
				response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
				return response, ctxErr3
			}
			response := &proto.ProfileResponse{Resparm1: ""}
			return response, nil
		}
	}

	response := &proto.ProfileResponse{Resparm1: "Follow user first"}
	return response, nil
}

func (*server) TweetService(ctx context.Context, request *proto.ProfileRequest) (*proto.ProfileResponse, error) {

	tweetContent := request.GetReqparm1()
	tweetUser := request.GetReqparm2()

	if tweetContent != "" {
		ctxErr := repository.SaveTweet(tweetUser,tweetContent,ctx)
		if ctxErr != nil{
			response := &proto.ProfileResponse{Resparm1: "Request timeout. Try again"}
			return response, ctxErr
		}
		response := &proto.ProfileResponse{Resparm1: ""}
		return response, nil
	} else {
		response := &proto.ProfileResponse{Resparm1: "Enter tweet content"}
		return response, nil
	}
}

func (*server) FeedService(ctx context.Context, request *proto.FeedRequest) (*proto.FeedResponse, error) {

	feedUser, _, ctxErr1 := repository.ReturnUser(request.GetReqparm1(),ctx)

	if ctxErr1 != nil{
		response := &proto.FeedResponse{Resparm1: "Request timeout. Try again",Resparm2: ""}
		return response, ctxErr1
	}

	feed := ""

	for i := 0 ; i < len(feedUser.Followers); i++{
		followUsername := feedUser.Followers[i]
		tweetList, ctxErr2 := repository.GetTweetList(followUsername,ctx)
		if ctxErr2 != nil{
			response := &proto.FeedResponse{Resparm1: "Request timeout. Try again",Resparm2: ""}
			return response, ctxErr2
		}
		if len(tweetList) != 0{
			feed = feed + GetTopFiveTweets(tweetList,followUsername)
		}
	}

	if feed != "" {
		feed = feed[:len(feed)-1]
		response := &proto.FeedResponse{Resparm1: "",Resparm2: feed}
		return response, nil
	} else {
		response := &proto.FeedResponse{Resparm1: "No feed",Resparm2: ""}
		return response, nil
	}
}

func GetTopFiveTweets(tweetList []string,followUsername string)(string){
	feed := ""
	if len(tweetList) > 5{
		tweetList = tweetList[len(tweetList)-5:]
	}
	for k := len(tweetList)-1;k >= 0; k-- {
		feed = feed + tweetList[k] + ","
	}

	if feed != ""{
		feed = feed[:len(feed)-1]
		feed = followUsername + "^" + feed + "$"
	}
	return feed

}

func (*server) UserListService(ctx context.Context, request *proto.FeedRequest) (*proto.FeedResponse, error) {

	userNameList, ctxErr1 := repository.GetUsers(ctx)

	if ctxErr1 != nil{
		response := &proto.FeedResponse{Resparm1: "Request timeout. Try again",Resparm2: ""}
		return response, ctxErr1
	}
	userNameList += "$"
	presentUser, _, ctxErr2 := repository.ReturnUser(request.GetReqparm1(),ctx)

	if ctxErr2 != nil{
		response := &proto.FeedResponse{Resparm1: "Request timeout. Try again",Resparm2: ""}
		return response, ctxErr2
	}

	for i:=0 ; i < len(presentUser.Followers); i++{
		userNameList += presentUser.Followers[i] + ","
	}
	if userNameList[len(userNameList)-1] == byte(','){
		userNameList = userNameList[:len(userNameList)-1]
	}

	response := &proto.FeedResponse{Resparm1: "",Resparm2: userNameList}
	return response, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	proto.RegisterTwitterServer(s, &server{})

	s.Serve(lis)
}
