package controllers

import (
	"gopkg.in/mgo.v2/bson"
	"regexp"
	"testapi/helpers"
	"testapi/models"
	"time"
)

var ps = models.NewPostResult()

type PostController struct {
	Model *models.Post
}

func NewPostController() *PostController {
	return &PostController{Model: &models.Post{}}
}

func (pc *PostController) GetPost(rt *helpers.Route) {
	id := regexp.MustCompile("/posts/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	Post, err := ps.FindOne(bson.M{"_id": bson.ObjectIdHex(id)})
	rt.Data["Post"] = Post

	if err != nil {
		//rt.Redirect("/admin/posts", 302)
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/404.html")
	} else {
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/post.html")
	}
}

func (pc *PostController) NewForm(rt *helpers.Route) {
	rt.Render("layout", false, "view/admin/layout.html", "view/admin/post.html")
}

func (pc *PostController) CreatePost(rt *helpers.Route) {
	rt.Request.ParseForm()
	data := rt.Request.Form

	Post := models.Post{
		Id:         bson.NewObjectId(),
		Title:      data.Get("title"),
		Preview:    data.Get("preview"),
		Text:       data.Get("text"),
		Created_at: time.Now(),
		Updated_at: time.Now()}
	ps.Collection.Insert(&Post)
	rt.Data["Post"] = Post

	if data.Get("action") == "save" {
		rt.Redirect("/admin/posts/"+Post.Id.Hex(), 302)
	} else {
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/post.html")
	}
}

func (pc *PostController) UpdatePost(rt *helpers.Route) {
	id := regexp.MustCompile("/posts/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]

	rt.Request.ParseForm()
	data := rt.Request.Form

	Post, err := ps.FindOne(bson.M{"_id": bson.ObjectIdHex(id)})
	if err == nil {
		Post.Title = data.Get("title")
		Post.Preview = data.Get("preview")
		Post.Code = data.Get("code")
		Post.Text = data.Get("text")
		Post.Updated_at = time.Now()
		if data.Get("active") == "1" {
			Post.Active = true
		} else {
			Post.Active = true
		}

		ps.Collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &Post)
	}
	rt.Data["Post"] = Post

	if data.Get("action") == "save" {
		rt.Redirect("/admin/posts/", 302)
	} else {
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/post.html")
	}
}

func (pc *PostController) DeletePost(rt *helpers.Route) {
	id := regexp.MustCompile("/posts/delete/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	ps.Collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	rt.Redirect("/admin/posts/", 302)
}

func (pc *PostController) IndexPost(rt *helpers.Route) {
	rt.Data["posts"] = ps.FindAll(bson.M{})
	rt.Render("layout", false, "view/admin/layout.html", "view/admin/posts.html")
}
