package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testapi/helpers"
)

type Post struct {
	Title      string
	Preview    string
	Text       string
	User       User
	Created_at string
	Updated_at string
}

type PostResult struct {
	DB         *helpers.DB
	Collection *mgo.Collection
	Result     interface{}
}

func (res *PostResult) GetCollection() *mgo.Collection {
	if res.Collection == nil {
		res.Collection = res.DB.GetCollection("posts")
	}

	return res.Collection
}

func (res *PostResult) FindOne(query bson.M) interface{} {
	res.GetCollection().Find(query).One(&res.Result)
	return res.Result
}
func (res *PostResult) FindAll(query bson.M) interface{} {
	res.GetCollection().Find(query).All(&res.Result)
	return res.Result
}
