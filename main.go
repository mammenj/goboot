package main

import (
	"github.com/mammenj/goboot/config"
	"github.com/mammenj/goboot/controllers"
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
	server :=  config.Server_port
	log.Printf("Started server %s .....",server)
	http.ListenAndServe(server, r)
}
