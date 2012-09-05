package main

import (
	"html/template"
	"path/filepath"
)

var cachedTemplates = map[string]*template.Template{}

var funcs = template.FuncMap{
	"reverse": reverse,
}

func T(name string) *template.Template {
	if t, ok := cachedTemplates[name]; ok {
		return t
	}

	t := template.New("_base.html").Funcs(funcs)

	t = template.Must(t.ParseFiles(
		"templates/_base.html",
		filepath.Join("templates", name),
	))
	cachedTemplates[name] = t

	return t
}
