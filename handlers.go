package main

import (
	"errors"
	"net/http"
)

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
	return T("index.html").Execute(w, map[string]interface{}{
		"entries": entries,
		"ctx":     ctx,
	})
}

func sign(w http.ResponseWriter, req *http.Request, ctx *Context) (err error) {
	//we need a user to sign to
	if ctx.User == nil {
		err = errors.New("Can't sign without being logged in")
		return
	}

	entry := NewEntry()
	entry.Name = ctx.User.Username
	entry.Message = req.FormValue("message")

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
	return T("login.html").Execute(w, map[string]interface{}{
		"ctx": ctx,
	})
}

func login(w http.ResponseWriter, req *http.Request, ctx *Context) error {
	username, password := req.FormValue("username"), req.FormValue("password")

	user, e := Login(ctx, username, password)
	if e != nil {
		ctx.Session.AddFlash("Invalid Username/Password")
		return loginForm(w, req, ctx)
	}

	//store the user id in the values and redirect to index
	ctx.Session.Values["user"] = user.ID
	http.Redirect(w, req, reverse("index"), http.StatusSeeOther)
	return nil
}

func logout(w http.ResponseWriter, req *http.Request, ctx *Context) error {
	delete(ctx.Session.Values, "user")
	http.Redirect(w, req, reverse("index"), http.StatusSeeOther)
	return nil
}
