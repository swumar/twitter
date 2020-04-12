package controller

import (
	// service "auth/service"
	// "log"
	"html/template"
	"net/http"

	"context"
	"auth/authpb"
	"google.golang.org/grpc"
	"log"
)


var opts = grpc.WithInsecure()
var cc, ccerr = grpc.Dial("localhost:50051", opts)


func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{}

	t, _ := template.ParseFiles("register.gtpl")

	if r.Method == "GET" {
		t.Execute(w, m)
		return 
    }else{

		if ccerr != nil {
			log.Fatal(ccerr)
		}

		defer cc.Close()

		client := authpb.NewRegisterServiceClient(cc)
		
		// log.Printf(r.Form["username"][1])
		// request := &authpb.RegisterRequest{Firstname: r.Form["firstname"][0], Lastname:r.Form["lastname"][0], Username:r.Form["username"][0], Password:r.Form["password"][0]}
		
		request := &authpb.RegisterRequest{Firstname: "Utkarsh", Lastname:"Prakash", Username:"up@gmail.com", Password:"up"}

		_, err := client.Register(context.Background(), request)

    	if err != nil {
			log.Printf("Receive Error Regiseter response => [%v]", err)
			m["Error"] = err
			log.Println(err)
			t.Execute(w, m)
			return
		}else{
			log.Println("User Registered succesfully")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	}
}

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
	
// 	t, _ := template.ParseFiles("login.gtpl")

// 	m := map[string]interface{}{}

// 	if r.Method == "GET" {
// 		t.Execute(w, nil)
// 		return 
// 	}else{

// 		errMsg := service.LoginService(w,r)

// 		if errMsg != "" {
// 			m["Error"] = errMsg
// 			log.Println(errMsg)
// 			t.Execute(w, m)
// 			return
// 		}else{
// 			log.Println("Login successful")
// 			http.Redirect(w, r, "/profile", http.StatusFound)
// 			return
// 		}
// 	}
// }

// func SignoutHandler(w http.ResponseWriter, r *http.Request) {

// 	t, _ := template.ParseFiles("login.gtpl")
// 	m := map[string]interface{}{}

// 	err := service.SignoutService(w,r)

// 	if err != nil {
// 		m["Error"] = "Please login to continue!"
// 		m["Success"] = nil
// 		log.Println("Please login to continue")
// 		t.Execute(w, m)
// 		return
// 	}else{
// 		log.Println("Logout succesfull")
// 		http.Redirect(w, r, "/login", http.StatusFound)
// 		return
// 	}
// }

