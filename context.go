package main

import (
	"code.google.com/p/gorilla/sessions"
	"labix.org/v2/mgo"
	"net/http"
)

type Context struct {
	Database *mgo.Database
	Session  *sessions.Session
}

func (c *Context) Close() {
	c.Database.Session.Close()
}

//C is a convenience function to return a collection from the context database.
func (c *Context) C(name string) *mgo.Collection {
	return c.Database.C(name)
}

func NewContext(req *http.Request) (*Context, error) {
	sess, err := store.Get(req, "gostbook")
	return &Context{
		Database: session.Clone().DB(database),
		Session:  sess,
	}, err
}
