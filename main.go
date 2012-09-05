package main

import (
	"code.google.com/p/gorilla/pat"
	"fmt"
	"labix.org/v2/mgo"
	"net/http"
	"os"
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
	session, err = mgo.Dial(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	database = session.DB("").Name

	router = pat.New()
	router.Add("GET", "/login", handler(loginForm)).Name("login")
	router.Add("GET", "/", handler(hello)).Name("index")
	router.Add("POST", "/sign", handler(sign)).Name("sign")

	if err = http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		panic(err)
	}
}
