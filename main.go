package main

import (
	//"./controllers"
	"fmt"
	"log"
	"net/http"
	"testapi/controllers"
	"testapi/helpers"
	//"regexp"
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

	http.Handle("/admin/theme/", http.StripPrefix("/admin/theme/", themeAdmin))
	http.Handle("/admin/public/", http.StripPrefix("/admin/public/", publicAdmin))
	http.Handle("/theme/", http.StripPrefix("/theme/", theme))
	http.Handle("/public/", http.StripPrefix("/public/", public))

	http.HandleFunc("/", startMainRoute)
	http.HandleFunc("/admin/", startAdminRoute)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func startAdminRoute(w http.ResponseWriter, r *http.Request) {

	topMenu.UpdateMenu(r.RequestURI)
	leftMenu.UpdateMenu(r.RequestURI)
	rt := helpers.Route{BaseURI: "/admin", Request: r, Response: w, Data: helpers.Data{"tmenu": topMenu, "lmenu": leftMenu}}
	pc := controllers.NewPostController()

	rt.Get("/posts/\\d+/?(|\\?.*)$", pc.GetPost)
	rt.Get("/posts/new/?(|\\?.*)$", pc.CreatePost)
	rt.Post("/posts/new/?(|\\?.*)$", pc.CreatePost)
	rt.Get("/posts/?(|\\?.*)$", pc.IndexPost)

	rt.Get("/(|\\?.*)$", func(rt *helpers.Route) {
		rt.Data["etts"] = "ers"
		rt.Render("layout", false, "view/admin/layout.html", "view/admin/index.html")
	})
	rt.Get("/login/?(|\\?.*)$", func(rt *helpers.Route) {
		rt.Render("layout", false, "view/admin/login.html")
	})

	rt.NotFound("layout", "view/admin/layout.html", "view/admin/404.html")
}

func startMainRoute(w http.ResponseWriter, r *http.Request) {

	rtMain := helpers.Route{BaseURI: "", Request: r, Response: w}
	fmt.Println("test")
	rtMain.Get("/(|\\?.*)$", func(rt *helpers.Route) {
		rt.Render("layout", false, "view/main/layout.html", "view/main/index.html")
	})
}
