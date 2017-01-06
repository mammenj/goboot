package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"

	"models"
)

var session *mgo.Session

type (
	UserController struct{}
)

func NewUserController() *UserController {
	s, err := mgo.Dial("mongodb://localhost")
	session = s
	if err != nil {
		panic(err)

	}
	return &UserController{}
}

/*
curl http://localhost:8001/user/hexid
*/
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	log.Printf("GetUser ID of user is >>>>> %s", id)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := session.DB("msa_DB").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	jsonU, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsonU)
}

/*
curl http://localhost:8001/users
*/
func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("List all Users  >>>>> ")
	us := []models.User{}
	if err := session.DB("msa_DB").C("users").Find(nil).All(&us); err != nil {
		log.Printf("GetUsers all Users  >>>>> %s ", err)
		w.WriteHeader(404)
		return
	}
	jsonUs, _ := json.Marshal(us)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsonUs)
}

/*
curl -XPOST -H 'Content-Type: application/json' -d
'{"name": "L John Mammen", "gender": "male", "age": 15}' http://localhost:8001/user
*/

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	log.Printf("Create User ID of user is >>>>> %s", u.Id)
	session.DB("msa_DB").C("users").Insert(u)
	jsonU, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", jsonU)
}

/*
curl -XDELETE http://localhost:8001/user/hexid
*/
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	log.Printf("RemoveUser ID of user is >>>>> %s", id)
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := session.DB("msa_DB").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(200)
}
