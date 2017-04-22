package main

import (
	"log"
	"net/http"
)

var topMenu = menu{
	{url: "/admin/", text: "Админка"},
	{url: "/admin/config/", text: "Настройки"},
	{url: "/admin/profile/", text: "Профиль"},
	{url: "/admin/help/", text: "Помощь"},
}

var leftMenu = menu{
	{url: "/admin/", text: "Главная"},
	{url: "/admin/posts/", text: "Статьи"},
	{url: "/admin/users/", text: "Пользователи"},
	{url: "/admin/ext/", text: "Дополнительно"},
}

func main() {
	public := http.FileServer(http.Dir("./public"))
	theme := http.FileServer(http.Dir("./themes/basic-dashboard"))

	http.Handle("/admin/theme/", http.StripPrefix("/admin/theme/", theme))
	http.Handle("/admin/public/", http.StripPrefix("/admin/public/", public))
	http.HandleFunc("/admin/", startRoute)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func startRoute(w http.ResponseWriter, r *http.Request) {
	topMenu.updateMenu(r.RequestURI)
	leftMenu.updateMenu(r.RequestURI)
	rt := Route{baseURI: "/admin", request: r, response: w, data: Data{"tmenu": topMenu, "lmenu": leftMenu}}

	rt.get("/posts/\\d+/?(|\\?.*)$", func(rt *Route) {
		rt.render("layout", false, "view/layout.html", "view/post.html")
	})
	rt.get("/posts/?(|\\?.*)$", func(rt *Route) {
		rt.data["etts"] = "ers"
		rt.render("layout", false, "view/layout.html", "view/posts.html")
	})
	rt.get("/(|\\?.*)$", func(rt *Route) {
		rt.data["etts"] = "ers"
		rt.render("layout", false, "view/layout.html", "view/index.html")
	})
	rt.get("/login/?(|\\?.*)$", func(rt *Route) {
		rt.render("layout", false, "view/login.html")
	})

	rt.notFound("layout", "view/layout.html", "view/404.html")
}
