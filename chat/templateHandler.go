package main

import (
	"html/template"
	"net/http"
	"path"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	tmpl     *template.Template
}

func (t *templateHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	t.once.Do(func() {
		t.tmpl = template.Must(template.ParseFiles(path.Join("templates", t.filename)))
	})
	t.tmpl.Execute(rw, req)
}
