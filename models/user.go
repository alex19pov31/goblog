package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testapi/helpers"
	"time"
)

type User struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	Login      string        `bson:"login"`
	Email      string        `bson:"email"`
	Password   string        `bson:"password"`
	Role       int           `bson:"role"`
	Active     bool          `bson:"active"`
	Created_at time.Time     `bson:"created_at"`
	Updated_at time.Time     `bson:"updated_at"`
}

// Проверка пароля
func (user *User) CheckPasswod(password, encPassword string) bool {

	return true
}

// Авторизация
func (user *User) Authorize(login, password string) bool {

	return true
}

// Возвращает зашифрованный пароль
func (user *User) GetPassword(password string) string {

	return password
}

func (user *User) GetId() string {
	return user.Id.Hex()
}

type UserResult struct {
	DB         *helpers.DB
	Collection *mgo.Collection
	Result     []User
}

func NewUserResult() *UserResult {
	db := helpers.DB{Host: "localhost", DBname: "testdb"}
	return &UserResult{DB: &db, Collection: db.GetCollection("users")}
}

func (res *UserResult) GetCollection() *mgo.Collection {
	if res.Collection == nil {
		res.Collection = res.DB.GetCollection("users")
	}

	return res.Collection
}

func (res *UserResult) FindOne(query bson.M) (User, error) {
	user := User{}
	err := res.GetCollection().Find(query).One(&user)
	return user, err
}
func (res *UserResult) FindAll(query bson.M) interface{} {
	res.GetCollection().Find(query).All(&res.Result)
	return res.Result
}
