package main

import (
	"html/template"
	"labix.org/v2/mgo"
	"net/http"
	"os"
)

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func hello(w http.ResponseWriter, req *http.Request) {
	//grab a clone of the session and close it when the
	//function returns
	s := session.Clone()
	defer s.Close()

	//set up the collection and query
	coll := s.DB("gostbook").C("entries")
	query := coll.Find(nil).Sort("-timestamp")

	//execute the query
	//TODO: add pagination :)
	var entries []Entry
	if err := query.All(&entries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//execute the template
	if err := index.Execute(w, entries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var session *mgo.Session

func main() {
	var err error
	session, err = mgo.Dial(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", hello)
	http.HandleFunc("/sign", sign)
	if err = http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
