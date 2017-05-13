package helpers

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"encoding/json"
)

type Data map[string]interface{}

type Route struct {
	BaseURI    string
	Checked    bool
	Data       Data
	Tpl        []string
	Pattern    string
	Request    *http.Request
	Response   http.ResponseWriter
	TemplatePath string
	BaseTemplate string
	BaseLayout   string
	credential Credential
}

type Credential struct {
	on bool
	isAllow bool
}

func (rt *Route) Route(method string, Pattern string, handler func(route *Route)) {
	if rt.Checked != true && rt.Request.Method == method && rt.CheckRegexp(Pattern) || method == "ANY" && rt.CheckRegexp(Pattern) {
		handler(rt)
	}
	rt.flashCredential()
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
	if rt.Checked == true || rt.credential.on && rt.credential.isAllow == false {
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

func (rt *Route) Show(tpl string, data interface{}, format string){
	if rt.TemplatePath == "" || rt.BaseTemplate == "" || rt.BaseLayout == "" {
		return
	}

	if(format == ".json") {
		rt.Response.Header().Set("Content-Type", "application/json; charset=utf-8")
		rt.Response.Write(GetJson(data))
	} else{
		rt.Render(rt.BaseLayout, data, rt.TemplatePath + rt.BaseTemplate, rt.TemplatePath + tpl)
	}
}

func (rt *Route) Credential(allow bool) *Route {
	rt.credential.on = true
	rt.credential.isAllow = allow

	return rt
}

func (rt *Route) flashCredential() {
	rt.credential.on = false
}

func GetJson(model interface{}) []byte {
	json, err := json.Marshal(model)
	if err != nil {
		return []byte{}
	}

	return json
}
