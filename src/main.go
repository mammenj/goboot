package main

import (
	"controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.GET("/users", uc.GetUsers)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.RemoveUser)
	http.ListenAndServe("localhost:8001", r)
}
