package controllers

import (
	"goblog/helpers"
	"goblog/models"
	"gopkg.in/mgo.v2/bson"
	"regexp"
)

type MainController struct {
}

func NewMainController() *MainController {
	return &MainController{}
}

func (mc *MainController) GetPost(rt *helpers.Route) {
	code := regexp.MustCompile("/post/([^/\\?]+)/?(|\\?.*)$").FindStringSubmatch(rt.Request.RequestURI)[1]
	Post, err := models.InitPost().FindOne(bson.M{"code": code})

	if err != nil {
		rt.Show("404.html", false, "html")
	}

	rt.Show("post.html", Post, "html")
}
