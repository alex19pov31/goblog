package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"strings"
)

var postCollection = dbConnect.GetCollection("posts")

type Post struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	Title      string        `bson:"title"`
	Code       string        `bson:"code"`
	Active     bool          `bson:"status"`
	Preview    string        `bson:"preview"`
	Text       string        `bson:"text"`
	User       User          `bson:"user"`
	Tags 	   []string      `bson:"tags"`
	Created_at time.Time     `bson:"created_at,omitempty"`
	Updated_at time.Time     `bson:"updated_at"`
}
type PostList []Post

func InitPost() Post {
	return Post{}
}

func (post *Post) GetId() string {
	return post.Id.Hex()
}

func (post *Post) Save() {
	post.Updated_at = time.Now()
	if post.Id == "" {
		post.Id = bson.NewObjectId()
		post.Created_at = post.Updated_at
		postCollection.Insert(&post)
	} else {
		postCollection.Update(bson.M{"_id": post.Id}, &post)
	}
}

func (post Post) GetTags(tags []string) string {
	return strings.Join(tags, ",")
}

func (post Post) FindOne(query bson.M) (Post, error) {
	err := postCollection.Find(query).One(&post)

	return post, err
}

func (post Post) FindById(id string) (Post, error) {
	err := postCollection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&post)

	return post, err
}

func (post Post) FindAll(query bson.M) PostList {
	var posts PostList
	postCollection.Find(query).All(&posts)

	return posts
}

func (post Post) Delete() error {
	err := postCollection.Remove(bson.M{"_id": post.Id})

	return err
}

func (post Post) DeleteByQuery(query bson.M) error {
	err := postCollection.Remove(query)

	return err
}

func (post Post) DeleteById(id string) error {
	err := postCollection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	return err
}