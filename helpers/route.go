package helpers

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

type Data map[string]interface{}

type Route struct {
	BaseURI  string
	Checked  bool
	Data     Data
	Tpl      []string
	Pattern  string
	Request  *http.Request
	Response http.ResponseWriter
}

func (rt *Route) Route(method string, Pattern string, handler func(route *Route)) {
	if rt.Request.Method == method && rt.CheckRegexp(Pattern) || method == "ANY" && rt.CheckRegexp(Pattern) {
		handler(rt)
	}
}
func (rt *Route) Get(Pattern string, handler func(route *Route)) {
	rt.Route("GET", Pattern, handler)
}
func (rt *Route) Post(Pattern string, handler func(route *Route)) {
	rt.Route("POST", Pattern, handler)
}
func (rt *Route) Put(Pattern string, handler func(route *Route)) {
	rt.Route("PUT", Pattern, handler)
}
func (rt *Route) Delete(Pattern string, handler func(route *Route)) {
	rt.Route("DELETE", Pattern, handler)
}

func (rt *Route) NotFound(layout string, Tpl ...string) {
	if rt.Checked != true {
		rt.Render(layout, false, Tpl...)
	}
}

func (rt *Route) Redirect(url string, code int) {
	http.Redirect(rt.Response, rt.Request, url, code)
}

func (rt *Route) CheckRegexp(Pattern string) bool {
	if rt.Checked == true {
		return false
	}
	if rt.BaseURI != "" && len(rt.BaseURI) > 0 {
		Pattern = rt.BaseURI + Pattern
	}
	check, err := regexp.MatchString(Pattern, rt.Request.RequestURI)
	if check && err == nil {
		rt.Checked = true
		return true
	}
	return false
}

func (rt *Route) Render(layout string, Data interface{}, Tpl ...string) {
	t, _ := template.ParseFiles(Tpl...)
	if Data == false {
		Data = rt.Data
	}
	if err := t.ExecuteTemplate(rt.Response, layout, Data); err != nil {
		log.Println(err)
	}
}
