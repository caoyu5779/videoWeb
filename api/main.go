package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)

	return router
}

func main()  {
	//每个goroutine 只有4k大小？
	r := RegisterHandlers()
	http.ListenAndServe(":8111", r)
}

//handler->validation{1: request , 2. user} -> business logic -> response
// 1. data model
// 2. error handling