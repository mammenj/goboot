package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mammenj/goboot/config"
	"github.com/mammenj/goboot/daos"
	"github.com/mammenj/goboot/models"
)

// MyUserController controller
type MyUserController struct {
	myuserDao daos.UserDao
}

// NewMyUserController creating controller
func NewMyUserController() *MyUserController {
	myconfig, err := config.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	myController := &MyUserController{}
	myController.myuserDao = daos.UserFactoryDao(myconfig.Engine)
	return myController
}

/*
GetUsers curl -GET http://localhost:8002/users
curl -GET http://localhost:8002/users
*/
func (uc MyUserController) GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("List all Users  >>>>> ")
	us, err := uc.myuserDao.GetAll()
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
CreateUser curl -XPOST -H 'Content-Type: application/json' -d '{"name": "L John Mammen", "gender": "male", "age": 15}' http://localhost:8002/user
curl -XPOST -H 'Content-Type: application/json' -d '{"name": "L John Mammen", "gender": "male", "age": 15}' http://localhost:8002/user
*/
func (uc MyUserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	err := uc.myuserDao.Create(&u)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Create User ID of user is >>>>> %d", u.Id)

	jsonU, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", jsonU)
}

// UpdateUser curl -H 'Content-Type: application/json' -H 'Accept: application/json' -X PUT -d '{"name": "L John Mammen", "gender": "male", "age": 15, "id":5}' http://localhost:8002/user
/*
curl -H 'Content-Type: application/json' -H 'Accept: application/json' -X PUT -d '{"name": "L John Mammen", "gender": "male", "age": 15, "id":5}' http://localhost:8002/user
*/

func (uc MyUserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	err := uc.myuserDao.Update(&u)
	if err != nil {
		log.Fatal(err)
		return
	}
	jsonU, _ := json.Marshal(u)
	log.Printf("Update User is >>>>> %s", jsonU)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", jsonU)
}

/*
RemoveUser curl -XDELETE http://localhost:8002/user/id
curl -XDELETE http://localhost:8002/user/id
*/
func (uc MyUserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	log.Printf("RemoveUser ID of user is >>>>> %d", id)

	err = uc.myuserDao.Delete(id)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Removed User ID of user is >>>>> %d", id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s %d", "Removed User", id)
}

/*
GetUser curl -GET http://localhost:8002/user/id
curl -GET http://localhost:8002/user/id
*/
func (uc MyUserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	log.Printf("GET user ID is >>>>> %d", id)
	user, err := uc.myuserDao.Get(id)
	jsonU, _ := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsonU)
}
