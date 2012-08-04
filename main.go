package main

import (
	"code.google.com/p/gorilla/pat"
	"fmt"
	"html/template"
	"labix.org/v2/mgo"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

var funcs = template.FuncMap{
	"reverse": reverse,
}

func parseTemplate(files ...string) *template.Template {
	//create a new template named after the first file in the list and add
	//the function map to it
	name := filepath.Base(files[0])
	t := template.New(name).Funcs(funcs)

	//parse the files into the template and panic on errors
	t = template.Must(t.ParseFiles(files...))
	return t
}

var index = parseTemplate(
	"templates/_base.html",
	"templates/index.html",
)

func reverse(name string, things ...interface{}) string {
	//convert the things to strings
	strs := make([]string, len(things))
	for i, th := range things {
		strs[i] = fmt.Sprint(th)
	}
	//grab the route
	u, err := router.GetRoute(name).URL(strs...)
	if err != nil {
		panic(err)
	}
	return u.Path
}

var session *mgo.Session
var database string
var router *pat.Router

func main() {
	var err error
	u := os.Getenv("DATABASE_URL")
	parsed, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	database = parsed.Path[1:]
	session, err = mgo.Dial(u)
	if err != nil {
		panic(err)
	}

	router = pat.New()
	router.Add("GET", "/", handler(hello)).Name("index")
	router.Add("POST", "/sign", handler(sign)).Name("sign")

	if err = http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		panic(err)
	}
}
