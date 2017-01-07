package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	"config"
	"daos"
	"models"
	"strconv"
)

type (
	MyUserController struct{}
)

var myuserDao daos.UserDao

func NewMyUserController() *MyUserController {
	myconfig, err := config.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	userDao := daos.UserFactoryDao(myconfig.Engine)
	myuserDao = userDao
	return &MyUserController{}
}

/*
curl -GET http://localhost:8002/users
*/
func (uc MyUserController) GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("List all Users  >>>>> ")
	us, err := myuserDao.GetAll()
	if err != nil {
		log.Fatal(err)
		return
	}
	jsonUs, _ := json.Marshal(us)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsonUs)
}

/*
curl -XPOST -H 'Content-Type: application/json' -d '{"name": "L John Mammen", "gender": "male", "age": 15}' http://localhost:8002/user
*/
func (uc MyUserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	err := myuserDao.Create(&u)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Create User ID of user is >>>>> %s", u.Id)
	jsonU, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", jsonU)
}

/*
curl -H 'Content-Type: application/json' -H 'Accept: application/json' -X PUT -d '{"name": "L John Mammen", "gender": "male", "age": 15, "id":5}' http://localhost:8002/user

 */

func (uc MyUserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	err := myuserDao.Update(&u)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Create User ID of user is >>>>> %s", u.Id)
	jsonU, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", jsonU)
}

/*
curl -XDELETE http://localhost:8002/user/id
*/
func (uc MyUserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	log.Printf("RemoveUser ID of user is >>>>> %s", id)

	err = myuserDao.Delete(id)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Removed User ID of user is >>>>> %s", id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s %n", "Removed User", id)
}

/*
curl -GET http://localhost:8002/user/id
*/
func (uc MyUserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	log.Printf("GET user ID is >>>>> %s", id)
	user, err := myuserDao.Get(id)
	jsonU, _ := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsonU)
}
