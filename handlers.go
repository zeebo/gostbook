package main

import "net/http"

func hello(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
	//set up the collection and query
	coll := ctx.C("entries")
	query := coll.Find(nil).Sort("-timestamp")

	//execute the query
	//TODO: add pagination :)
	var entries []Entry
	if err = query.All(&entries); err != nil {
		return
	}

	//execute the template
	return T("index.html").Execute(w, entries)
}

func sign(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
	entry := NewEntry()
	entry.Name = req.FormValue("name")
	entry.Message = req.FormValue("message")

	if entry.Name == "" {
		entry.Name = "Some dummy who forgot a name"
	}
	if entry.Message == "" {
		entry.Message = "Some dummy who forgot a message."
	}

	coll := ctx.C("entries")
	if err = coll.Insert(entry); err != nil {
		return
	}

	http.Redirect(w, req, reverse("index"), http.StatusSeeOther)
	return
}

func loginForm(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
	return T("login.html").Execute(w, nil)
}
