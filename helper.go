package main

import (
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func render(tpl string, data interface{}, w http.ResponseWriter) {
	t, _ := template.ParseFiles("view/layout.html", "view/"+tpl+".html")
	if err := t.ExecuteTemplate(w, "layout", data); err != nil {
		log.Println(err)
	}
}

func checkAuth(user, password string, w http.ResponseWriter, r *http.Request) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	return pair[0] == user && pair[1] == password
}

func setAuth(user, password string, w http.ResponseWriter, r *http.Request) bool {
	if checkAuth(user, password, w, r) == false {
		w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
		return false
	}
	return true
}
