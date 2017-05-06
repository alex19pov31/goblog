package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"golang.org/x/crypto/bcrypt"
)

var userCollection = dbConnect.GetCollection("users")

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
type UserList []User

func InitUser() User{
	return User{}
}

// Проверка пароля
func (user *User) CheckPasswod(password, encPassword string) bool {
	password, err := user.GetPassword(password)
	return password == encPassword && err == nil
}

// Авторизация
func (user *User) Authorize(login, password string) bool {

	return true
}

// Возвращает зашифрованный пароль
func (user *User) GetPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (user *User) GetId() string {
	return user.Id.Hex()
}

func (user *User) Save() {
	user.Updated_at = time.Now()
	if user.Id == "" {
		user.Id = bson.NewObjectId()
		user.Created_at = user.Updated_at
		userCollection.Insert(&user)
	} else {
		userCollection.Update(bson.M{"_id": user.Id}, &user)
	}
}

func (user User) FindOne(query bson.M) (User, error) {
	err := userCollection.Find(query).One(&user)

	return user, err
}

func (user User) FindById(id string) (User, error)  {
	err := userCollection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)

	return user, err
}

func (user User) FindAll(query bson.M) UserList {
	var users UserList
	userCollection.Find(query).All(&users)

	return users
}

func (user User) Delete() error {
	err := userCollection.Remove(bson.M{"_id": user.Id})

	return  err
}

func (user User) DeleteByQuery(query bson.M) error {
	err := userCollection.Remove(query)

	return err
}

func (user User) DeleteById(id string) error {
	err := userCollection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	return err
}