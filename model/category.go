package model

type Category struct {
	ID   string `json:"ID" bson:"_id"`
	Name string `json:"name" bson:"name"`
}
