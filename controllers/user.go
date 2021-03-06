package controllers

import (
	"goblog/helpers"
	"goblog/models"
	"gopkg.in/mgo.v2/bson"
	"net/url"
	"regexp"
	"strconv"
)

type UserController struct {
	Model *models.User
}

func NewUserController() *UserController {
	return &UserController{Model: &models.User{}}
}

func (pc *UserController) Get(rt *helpers.Route) {
	id := regexp.MustCompile("/users/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	user, err := models.InitUser().FindById(id)
	rt.Data["User"] = user

	if err != nil {
		rt.Show("404.html", false, "")
	} else {
		rt.Show("user.html", false, "")
	}
}

func (pc *UserController) New(rt *helpers.Route) {
	rt.Show("user.html", false, "")
}

func (pc *UserController) Create(rt *helpers.Route) {
	rt.Request.ParseForm()
	data := rt.Request.Form
	active, _ := strconv.ParseBool(data.Get("active"))
	role, _ := strconv.Atoi(data.Get("role"))

	User := models.User{
		Login:  data.Get("login"),
		Email:  data.Get("email"),
		Role:   role,
		Active: active}
	pc.setPassword(data, &User)

	User.Save()
	rt.Data["User"] = User

	if data.Get("action") == "save" {
		rt.Redirect("/admin/users/"+User.Id.Hex(), 302)
	} else {
		rt.Show("user.html", false, "")
	}
}

func (pc *UserController) Update(rt *helpers.Route) {
	id := regexp.MustCompile("/users/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	rt.Request.ParseForm()
	data := rt.Request.Form
	active, _ := strconv.ParseBool(data.Get("active"))
	role, _ := strconv.Atoi(data.Get("role"))
	User, userErr := models.InitUser().FindById(id)

	helpers.CheckErrors(userErr)

	User.Login = data.Get("login")
	User.Email = data.Get("email")
	User.Role = role
	User.Active = active
	pc.setPassword(data, &User)

	User.Save()

	rt.Data["User"] = User

	if data.Get("action") == "save" {
		rt.Redirect("/admin/users/", 302)
	} else {
		rt.Show("user.html", false, "")
	}
}

func (pc *UserController) Delete(rt *helpers.Route) {
	id := regexp.MustCompile("/users/delete/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	models.InitUser().DeleteById(id)
	rt.Redirect("/admin/users/", 302)
}

func (pc *UserController) Active(rt *helpers.Route) {
	id := regexp.MustCompile("/users/active/([^/\\?]{24})/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	User, userErr := models.InitUser().FindById(id)
	helpers.CheckErrors(userErr)

	User.Active = !User.Active
	User.Save()

	rt.Redirect("/admin/users/", 302)
}

func (pc *UserController) Index(rt *helpers.Route) {
	rt.Data["users"] = models.InitUser().FindAll(bson.M{})
	rt.Show("users.html", false, "")
}

func (pc *UserController) Login(rt *helpers.Route) {
	rt.Request.ParseForm()
	data := rt.Request.Form

	models.Login(rt.Response, data.Get("login"), data.Get("password"))

	rt.Redirect("/admin/", 302)
}

func (pc *UserController) Logout(rt *helpers.Route) {
	models.Logout(rt.Response)
	rt.Redirect("/admin/", 302)
}

func (pc *UserController) setPassword(data url.Values, user *models.User) {
	password1 := data.Get("password")
	password2 := data.Get("repeatPassword")

	if password1 != "" && password1 == password2 && len(password1) > 6 {
		password, passErr := models.GetPassword(password1)
		if passErr == nil {
			user.Password = password
		}
	}
}
