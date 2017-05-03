package controllers

import (
	"gopkg.in/mgo.v2/bson"
	"regexp"
	"testapi/helpers"
	"testapi/models"
	"time"
)

var ur = models.NewUserResult()

type UserController struct {
	Model *models.User
}

func NewUserController() *UserController {
	return &UserController{Model: &models.User{}}
}

func (pc *UserController) Get(rt *helpers.Route) {
	id := regexp.MustCompile("/posts/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	Post, err := ur.FindOne(bson.M{"_id": bson.ObjectIdHex(id)})
	rt.Data["Post"] = Post

	if err != nil {
		//rt.Redirect("/admin/posts", 302)
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/404.html")
	} else {
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/user.html")
	}
}

func (pc *UserController) New(rt *helpers.Route) {
	rt.Render("layout", false, "view/admin/layout.html", "view/admin/user.html")
}

func (pc *UserController) Create(rt *helpers.Route) {
	rt.Request.ParseForm()
	data := rt.Request.Form

	User := models.User{
		Id:         bson.NewObjectId(),
		Login:      data.Get("login"),
		Email:      data.Get("email"),
		Role:       data.Get("role"),
		Password:   data.Get("password1"),
		Active:     data.Get("active"),
		Created_at: time.Now(),
		Updated_at: time.Now()}
	ur.Collection.Insert(&User)
	rt.Data["User"] = User

	if data.Get("action") == "save" {
		rt.Redirect("/admin/users/"+User.Id.Hex(), 302)
	} else {
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/user.html")
	}
}

func (pc *UserController) Update(rt *helpers.Route) {
	id := regexp.MustCompile("/users/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]

	rt.Request.ParseForm()
	data := rt.Request.Form
	Post, err := ps.FindOne(bson.M{"_id": bson.ObjectIdHex(id)})
	if err == nil {
		User.Login = data.Get("login")
		User.Email = data.Get("email")
		User.Role = data.Get("role")
		User.Password = data.Get("password1")
		User.Active = data.Get("active")
		User.Updated_at = time.Now()
		ur.Collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &User)
	}
	rt.Data["User"] = User

	if data.Get("action") == "save" {
		rt.Redirect("/admin/users/", 302)
	} else {
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/user.html")
	}
}

func (pc *UserController) Delete(rt *helpers.Route) {
	id := regexp.MustCompile("/users/delete/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	ur.Collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	rt.Redirect("/admin/users/", 302)
}

func (pc *UserController) Active(rt *helpers.Route) {
	id := regexp.MustCompile("/posts/users/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	Post, err := ps.FindOne(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		Post.Active = true
		User.Updated_at = time.Now()
		ur.Collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &User)
	}
	ur.Collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	rt.Redirect("/admin/posts/", 302)
}

func (pc *UserController) Index(rt *helpers.Route) {
	rt.Data["users"] = ur.FindAll(bson.M{})
	rt.Render("layout", false, "view/admin/layout.html", "view/admin/users.html")
}
