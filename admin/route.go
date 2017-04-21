package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

type Route struct {
	baseURI  string
	checked  bool
	data     interface{}
	tpl      []string
	request  *http.Request
	response http.ResponseWriter
}

func (rt *Route) route(method string, pattern string, handler func(route *Route)) {
	if rt.request.Method == method && rt.checkRegexp(pattern) || method == "ANY" && rt.checkRegexp(pattern) {
		handler(rt)
	}
}
func (rt *Route) get(pattern string, handler func(route *Route)) {
	rt.route("GET", pattern, handler)
}
func (rt *Route) post(pattern string, handler func(route *Route)) {
	rt.route("POST", pattern, handler)
}
func (rt *Route) put(pattern string, handler func(route *Route)) {
	rt.route("PUT", pattern, handler)
}
func (rt *Route) delete(pattern string, handler func(route *Route)) {
	rt.route("DELETE", pattern, handler)
}

func (rt *Route) notFound(layout string, tpl ...string) {
	if rt.checked != true {
		rt.render(layout, false, tpl...)
	}
}

func (rt *Route) checkRegexp(pattern string) bool {
	if rt.checked == true {
		return false
	}
	if len(rt.baseURI) > 0 {
		pattern = rt.baseURI + pattern
	}
	check, err := regexp.MatchString(pattern, rt.request.RequestURI)
	if check && err == nil {
		rt.checked = true
		return true
	}
	return false
}

func (rt *Route) render(layout string, data interface{}, tpl ...string) {
	t, _ := template.ParseFiles(tpl...)
	if err := t.ExecuteTemplate(rt.response, layout, data); err != nil {
		log.Println(err)
	}
}
