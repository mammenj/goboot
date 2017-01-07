package main

import (
	"config"
	"controllers"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	config, err := config.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}
	r := httprouter.New()
	uc := controllers.NewMyUserController()
	r.GET("/user/:id", uc.GetUser)
	r.GET("/users", uc.GetUsers)
	r.POST("/user", uc.CreateUser)
	r.PUT("/user", uc.UpdateUser)
	r.DELETE("/user/:id", uc.RemoveUser)
	server := "localhost:" + config.Serverport
	http.ListenAndServe(server, r)
}
