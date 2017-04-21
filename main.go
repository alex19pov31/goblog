package main

import (
	"log"
	"net/http"
)

func main() {
	public := http.FileServer(http.Dir("./public"))
	theme := http.FileServer(http.Dir("./themes/startbootstrap-blog-post"))

	http.Handle("/theme/", http.StripPrefix("/theme/", theme))
	http.Handle("/public/", http.StripPrefix("/public/", public))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/about", aboutHandler)
	//http.HandleFunc("/admin", adminHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"H": "test123", "G": "test", "t": "tttest"}
	render("index", data, w)

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"H": "test123", "G": "test", "t": "tttest"}
	render("about", data, w)

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	render("test", map[string]interface{}{"H": "test123", "G": "test", "t": "tttest"}, w)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	if setAuth("test", "test", w, r) {
		render("admin", "", w)
	}
}
