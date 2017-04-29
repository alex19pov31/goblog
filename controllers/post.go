package controllers

import (
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
	//id := regexp.MustCompile("/posts/(\\d+)/?(|\\?.*)$").FindStringSubmatch(rt.request.RequestURI)[1]

	db := helpers.DB{Host: "localhost", DBname: "testdb"}
	db.GetCollection("posts")

	rt.Render("layout", false, "view/admin/layout.html", "view/admin/post.html")
}

func (pc *PostController) NewForm(rt *helpers.Route) {
	rt.Render("layout", false, "view/admin/layout.html", "view/admin/post.html")
}

func (pc *PostController) CreatePost(rt *helpers.Route) {
	rt.Request.ParseForm()
	data := rt.Request.Form

	Post := models.Post{
		Title:      data.Get("title"),
		Preview:    data.Get("preview"),
		Text:       data.Get("text"),
		Created_at: time.Now(),
		Updated_at: time.Now()}

	ps.Collection.Insert(&Post)
	rt.Data["Post"] = Post

	rt.Render("layout", false, "view/admin/layout.html", "view/admin/post.html")
}

func (pc *PostController) UpdatePost(rt *helpers.Route) {
	//var result interface{}

	//id := 1
	//pc.collection.Find(bson.M{"_id": id}).One(&result)
}

func (pc *PostController) DeletePost(rt *helpers.Route) {

}

func (pc *PostController) IndexPost(rt *helpers.Route) {
	rt.Data["etts"] = "ers"
	rt.Render("layout", false, "view/admin/layout.html", "view/admin/posts.html")
}
