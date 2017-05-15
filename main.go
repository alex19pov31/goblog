package main

import (
	"goblog/controllers"
	"goblog/helpers"
	"goblog/models"
	"log"
	"net/http"
)

var isAuthorize bool

var topMenu = helpers.Menu{
	{URL: "/admin/", Text: "Админка"},
	{URL: "/admin/config/", Text: "Настройки"},
	{URL: "/admin/profile/", Text: "Профиль"},
	{URL: "/admin/help/", Text: "Помощь"},
}

var leftMenu = helpers.Menu{
	{URL: "/admin/", Text: "Главная"},
	{URL: "/admin/posts/", Text: "Статьи"},
	{URL: "/admin/users/", Text: "Пользователи"},
	{URL: "/admin/ext/", Text: "Дополнительно"},
}

func main() {
	publicAdmin := http.FileServer(http.Dir("./public"))
	themeAdmin := http.FileServer(http.Dir("./themes/basic-dashboard"))
	public := http.FileServer(http.Dir("./public"))
	theme := http.FileServer(http.Dir("./themes/startbootstrap-blog"))

	go http.Handle("/admin/theme/", http.StripPrefix("/admin/theme/", themeAdmin))
	go http.Handle("/admin/public/", http.StripPrefix("/admin/public/", publicAdmin))
	go http.Handle("/theme/", http.StripPrefix("/theme/", theme))
	go http.Handle("/public/", http.StripPrefix("/public/", public))

	go http.HandleFunc("/", startMainRoute)
	go http.HandleFunc("/admin/", startAdminRoute)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func startAdminRoute(w http.ResponseWriter, r *http.Request) {
	Authorize(w, r)

	topMenu.UpdateMenu(r.RequestURI)
	leftMenu.UpdateMenu(r.RequestURI)
	route := helpers.Route{
		BaseURI:      "/admin",
		Request:      r,
		Response:     w,
		TemplatePath: "view/admin/",
		BaseTemplate: "layout.html",
		BaseLayout:   "layout",
		Data:         helpers.Data{"tmenu": topMenu, "lmenu": leftMenu},
	}

	defer route.NotFound()

	postsRoute(&route)
	usersRoute(&route)

	route.Credential(models.CurUser.Role == 1).Get("/(|\\?.*)$", func(route *helpers.Route) {
		route.Show("index.html", false, "html")
		return
	})

	route.Get("/login/?(|\\?.*)$", func(route *helpers.Route) {
		if isAuthorize == false {
			route.Render("layout", false, "view/admin/login.html")
		} else {
			route.Redirect("/admin/", 302)
		}
	})
}

func postsRoute(route *helpers.Route) {
	cnt := controllers.NewPostController()
	allowRule := models.CurUser.Role == 1

	route.Credential(allowRule).Get("/posts/publish/[^/\\?]{24}/?(|\\?.*)$", cnt.Active)
	route.Credential(allowRule).Get("/posts/delete/[^/\\?]{24}/?(|\\?.*)$", cnt.Delete)
	route.Credential(allowRule).Get("/posts/new/?(|\\?.*)$", cnt.New)
	route.Credential(allowRule).Post("/posts/new/?(|\\?.*)$", cnt.Create)
	route.Credential(allowRule).Get("/posts/[^/\\?]{24}(/?|\\.json|/?\\?.*)$", cnt.Get)
	route.Credential(allowRule).Post("/posts/[^/\\?]{24}/?(|\\?.*)$", cnt.Update)
	route.Credential(allowRule).Get("/posts(/?|\\.json|/?\\?.*)$", cnt.Index)
}

func usersRoute(route *helpers.Route) {
	cnt := controllers.NewUserController()
	allowRule := models.CurUser.Role == 1

	route.Post("/login/?(|\\?.*)$", cnt.Login)
	route.Get("/logout/?(|\\?.*)$", cnt.Logout)
	route.Credential(allowRule).Get("/users/new?(|\\?.*)$", cnt.New)
	route.Credential(allowRule).Post("/users/new?(|\\?.*)$", cnt.Create)
	route.Credential(allowRule).Get("/users/[^/\\?]{24}/?(|\\?.*)$", cnt.Get)
	route.Credential(allowRule).Post("/users/[^/\\?]{24}/?(|\\?.*)$", cnt.Update)
	route.Credential(allowRule).Get("/users/delete/[^/\\?]{24}/?(|\\?.*)$", cnt.Delete)
	route.Credential(allowRule).Get("/users/active/[^/\\?]{24}/?(|\\?.*)$", cnt.Active)
	route.Credential(allowRule).Get("/users/?(|\\?.*)$", cnt.Index)
}

func startMainRoute(w http.ResponseWriter, r *http.Request) {

	route := helpers.Route{
		BaseURI:      "/",
		Request:      r,
		Response:     w,
		TemplatePath: "view/main/",
		BaseTemplate: "layout.html",
		BaseLayout:   "layout",
		Data:         helpers.Data{},
	}

	defer route.NotFound()

	route.Get("/?(|\\?.*)$", func(rt *helpers.Route) {
		route.Show("index.html", false, "html")
	})
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	user, isAuth := models.AuthorizeByCookie(r)
	isAuthorize = isAuth
	user.Password = ""
	models.CurUser = user

	if isAuth == false && r.RequestURI != "/admin/login" {
		http.Redirect(w, r, "/admin/login", 302)
	}
}
