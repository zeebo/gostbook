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

func NewContext(req *http.Request) (*Context, error) {
	return &Context{
		Database: session.Clone().DB(database),
	}, nil
}
