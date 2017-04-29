package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testapi/helpers"
	"time"
)

type Post struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	Title      string        `bson:"title"`
	Preview    string        `bson:"preview"`
	Text       string        `bson:"text"`
	User       User          `bson:"user"`
	Created_at time.Time     `bson:"created_at"`
	Updated_at time.Time     `bson:"updated_at"`
}

type PostResult struct {
	DB         *helpers.DB
	Collection *mgo.Collection
	Result     interface{}
}

func NewPostResult() *PostResult {
	db := helpers.DB{Host: "localhost", DBname: "testdb"}
	return &PostResult{DB: &db, Collection: db.GetCollection("posts")}
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
