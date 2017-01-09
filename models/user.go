package models

type (
	User struct {
		Id     int    `json:"id" bson:"_id"`
		Name   string `json:"name" bson:"name"`
		Gender string `json:"gender" bson:"gender"`
		Age    int    `json:"age" bson:"age"`
	}
)
