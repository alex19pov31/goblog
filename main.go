package main

import (
	"log"
	"net/http"
	"testapi/controllers"
	"testapi/helpers"
)

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
	theme := http.FileServer(http.Dir("./themes/startbootstrap-blog-post"))

	go http.Handle("/admin/theme/", http.StripPrefix("/admin/theme/", themeAdmin))
	go http.Handle("/admin/public/", http.StripPrefix("/admin/public/", publicAdmin))
	go http.Handle("/theme/", http.StripPrefix("/theme/", theme))
	go http.Handle("/public/", http.StripPrefix("/public/", public))

	go http.HandleFunc("/", startMainRoute)
	go http.HandleFunc("/admin/", startAdminRoute)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func startAdminRoute(w http.ResponseWriter, r *http.Request) {
	topMenu.UpdateMenu(r.RequestURI)
	leftMenu.UpdateMenu(r.RequestURI)
	route := helpers.Route{BaseURI: "/admin", Request: r, Response: w, Data: helpers.Data{"tmenu": topMenu, "lmenu": leftMenu}}

	postsRoute(&route)
	usersRoute(&route)

	route.Get("/(|\\?.*)$", func(route *helpers.Route) {
		route.Render("layout", false, "view/admin/layout.html", "view/admin/index.html")
	})
	route.Get("/login/?(|\\?.*)$", func(route *helpers.Route) {
		route.Render("layout", false, "view/admin/login.html")
	})

	route.NotFound("layout", "view/admin/layout.html", "view/admin/404.html")
}

func postsRoute(route *helpers.Route) {
	cnt := controllers.NewPostController()

	route.Get("/posts/delete/[^/\\?]{24}/?(|\\?.*)$", cnt.Delete)
	route.Get("/posts/new/?(|\\?.*)$", cnt.New)
	route.Post("/posts/new/?(|\\?.*)$", cnt.Create)
	route.Get("/posts/[^/\\?]{24}/?(|\\?.*)$", cnt.Get)
	route.Post("/posts/[^/\\?]{24}/?(|\\?.*)$", cnt.Update)
	route.Get("/posts/?(|\\?.*)$", cnt.Index)
}

func usersRoute(route *helpers.Route) {
	cnt := controllers.NewUserController()

	route.Get("/users/new?(|\\?.*)$", cnt.New)
	route.Post("/users/new?(|\\?.*)$", cnt.Create)
	route.Get("/users/[^/\\?]{24}/?(|\\?.*)$", cnt.Get)
	route.Post("/users/[^/\\?]{24}/?(|\\?.*)$", cnt.Update)
	route.Get("/users/delete/[^/\\?]{24}/?(|\\?.*)$", cnt.Delete)
	route.Get("/users/active/[^/\\?]{24}/?(|\\?.*)$", cnt.Active)
	route.Get("/users/?(|\\?.*)$", cnt.Index)
}

func startMainRoute(w http.ResponseWriter, r *http.Request) {

	rtMain := helpers.Route{BaseURI: "", Request: r, Response: w}
	rtMain.Get("/(|\\?.*)$", func(rt *helpers.Route) {
		rt.Render("layout", false, "view/main/layout.html", "view/main/index.html")
	})
}
