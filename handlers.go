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
	err = index.Execute(w, entries)
	return
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

func login(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
	//grab the username and password from the form
	username, password := req.FormValue("username"), req.FormValue("password")

	//log in the user
	user, err := Login(ctx, username, password)

	//what to do now? if there was an error we want to present the form again
	//with some error message.

	//where do we store the user if the login was valid?

	//answer: sessions!
	_ = user
	return
}
