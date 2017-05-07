package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"encoding/json"
	"crypto"
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

type Auth struct {
	login string	`json:"login"`
	token string	`json:"token"`
	ttl time.Time	`ttl:"ttl"`
}

func InitUser() User{
	return User{}
}

// Проверка пароля
func CheckPasswod(password, encPassword string) bool {
	password, err := GetPassword(password)
	return password == encPassword && err == nil
}

// Авторизация
func Login(w http.ResponseWriter, login, password string) (User, bool) {
	user, err := InitUser().FindOne(bson.M{"login": login})
	check := err == nil && CheckPasswod(password, user.Password)

	if check {
		setCookie(w, user)
		return user, true
	}

	return User{}, false
}

func AuthorizeByCookie(r http.Request) (User, bool){
	auth, authErr := getCookie(r)
	user, userErr := InitUser().FindOne(bson.M{"login": auth.login});
	userToken := getToken(user, auth.ttl)

	if auth.token == userToken && authErr == nil && userErr == nil {
		return user, true
	}

	return User{}, false
}

func setCookie(w http.ResponseWriter, user User) {
	ttl := time.Now().Local().Add(time.Hour * 24)
	tk := getToken(user, ttl)
	auth := Auth{
		login: user.Login,
		token: tk,
		ttl: ttl,
	}

	strJSON, _ := json.Marshal(auth)
	cookie := http.Cookie{
		Name:"gblock",
		Value: string(strJSON),
		Path: "/",
		Expires: ttl,
	}

	http.SetCookie(w, &cookie)
}

func getToken(user User, ttl time.Time) string {
	return string(crypto.MD5.New().Sum([]byte(user.Login + user.Password + ttl.String())))
}

func getCookie(r http.Request) (Auth, error){
	cookie, err := r.Cookie("gblock")
	auth := Auth{}
	if err == nil {
		json.Unmarshal([]byte(cookie.Value), &auth)
	}

	return auth, err
}

// Возвращает зашифрованный пароль
func GetPassword(password string) (string, error) {
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