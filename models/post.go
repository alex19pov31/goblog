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
	Created_at time.Time     `bson:"created_at,omitempty"`
	Updated_at time.Time     `bson:"updated_at"`
}

func (post *Post) GetId() string {
	return post.Id.Hex()
}

type PostResult struct {
	DB         *helpers.DB
	Collection *mgo.Collection
	Result     []Post
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

func (res *PostResult) FindOne(query bson.M) (Post, error) {
	post := Post{}
	err := res.GetCollection().Find(query).One(&post)
	return post, err
}
func (res *PostResult) FindAll(query bson.M) interface{} {
	res.GetCollection().Find(query).All(&res.Result)
	return res.Result
}
