package main

import (
	"html/template"
	"labix.org/v2/mgo"
	"net/http"
	"net/url"
	"os"
)

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

var session *mgo.Session
var database string

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

	http.Handle("/", handler(hello))
	http.Handle("/sign", handler(sign))
	if err = http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
