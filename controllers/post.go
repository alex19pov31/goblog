package controllers

import (
	"gopkg.in/mgo.v2/bson"
	"regexp"
	"log"
	"strconv"
	"../helpers"
	"../models"
)

type PostController struct {
	Model *models.Post
}

func NewPostController() *PostController {
	return &PostController{Model: &models.Post{}}
}

func (pc *PostController) Get(rt *helpers.Route) {
	id := regexp.MustCompile("/posts/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	Post, err := models.InitPost().FindById(id)
	rt.Data["Post"] = Post

	if err != nil {
		//rt.Redirect("/admin/posts", 302)
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/404.html")
	} else {
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/post.html")
	}
}

func (pc *PostController) New(rt *helpers.Route) {
	rt.Render("layout", false, "view/admin/layout.html", "view/admin/post.html")
}

func (pc *PostController) Create(rt *helpers.Route) {
	rt.Request.ParseForm()
	data := rt.Request.Form

	active, _ := strconv.ParseBool(data.Get("active"))
	Post := models.Post{
		Title:      data.Get("title"),
		Preview:    data.Get("preview"),
		Code:       data.Get("code"),
		Text:       data.Get("text"),
		Active:     active}
	Post.Save()
	rt.Data["Post"] = Post

	if data.Get("action") == "save" {
		rt.Redirect("/admin/posts/"+Post.Id.Hex(), 302)
	} else {
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/post.html")
	}
}

func (pc *PostController) Update(rt *helpers.Route) {
	id := regexp.MustCompile("/posts/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]

	rt.Request.ParseForm()
	data := rt.Request.Form

	Post, postErr := models.InitPost().FindById(id)
	helpers.CheckErrors(postErr)

	Post.Title = data.Get("title")
	Post.Preview = data.Get("preview")
	Post.Code = data.Get("code")
	Post.Text = data.Get("text")
	Post.Active, _ = strconv.ParseBool(data.Get("active"))
	Post.Save()

	rt.Data["Post"] = Post

	if data.Get("action") == "save" {
		rt.Redirect("/admin/posts/", 302)
	} else {
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/post.html")
	}
}

func (pc *PostController) Delete(rt *helpers.Route) {
	id := regexp.MustCompile("/posts/delete/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	err := models.InitPost().DeleteById(id)
	log.Println(err)
	rt.Redirect("/admin/posts/", 302)
}

func (pc *PostController) Index(rt *helpers.Route) {
	rt.Data["posts"] = models.InitPost().FindAll(bson.M{})
	rt.Render("layout", false, "view/admin/layout.html", "view/admin/posts.html")
}


func (pc *PostController) Active(rt *helpers.Route) {
	id := regexp.MustCompile("/posts/publish/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	Post, postErr := models.InitPost().FindById(id)
	helpers.CheckErrors(postErr)

	Post.Active = !Post.Active
	Post.Save()

	rt.Redirect("/admin/posts/", 302)
}