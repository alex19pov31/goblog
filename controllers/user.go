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

	Post := models.Post{
		Title:      data.Get("title"),
		Preview:    data.Get("preview"),
		Text:       data.Get("text"),
		Created_at: time.Now(),
		Updated_at: time.Now()}
	ur.Collection.Insert(&Post)
	rt.Data["Post"] = Post

	if data.Get("action") == "save" {
		rt.Redirect("/admin/posts/", 302)
	} else {
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/post.html")
	}
}

func (pc *UserController) Update(rt *helpers.Route) {
	id := regexp.MustCompile("/posts/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]

	rt.Request.ParseForm()
	data := rt.Request.Form
	Post := models.Post{
		Title:      data.Get("title"),
		Preview:    data.Get("preview"),
		Text:       data.Get("text"),
		Updated_at: time.Now()}
	ur.Collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &Post)
	rt.Data["Post"] = Post

	if data.Get("action") == "save" {
		rt.Redirect("/admin/posts/", 302)
	} else {
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/post.html")
	}
}

func (pc *UserController) Delete(rt *helpers.Route) {
	id := regexp.MustCompile("/posts/delete/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	ur.Collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	rt.Redirect("/admin/posts/", 302)
}

func (pc *UserController) Active(rt *helpers.Route) {
	id := regexp.MustCompile("/posts/delete/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	ur.Collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	rt.Redirect("/admin/posts/", 302)
}

func (pc *UserController) Index(rt *helpers.Route) {
	rt.Data["users"] = ur.FindAll(bson.M{})
	rt.Render("layout", false, "view/admin/layout.html", "view/admin/users.html")
}
