package main

import (
	"labix.org/v2/mgo"
	"net/http"
)

type Context struct {
	Database *mgo.Database
}

func (c *Context) Close() {
	c.Database.Session.Close()
}

//C is a convenience function to return a collection from the context database.
func (c *Context) C(name string) *mgo.Collection {
	return c.Database.C(name)
}

func NewContext(req *http.Request) (*Context, error) {
	return &Context{
		Database: session.Clone().DB(database),
	}, nil
}
