package main

import (
	"html/template"
	"net/http"
	"path"
	"sync"

	"github.com/stretchr/objx"
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
	data := map[string]interface{}{
		"Host": req.Host,
	}
	if authCookie, err := req.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.tmpl.Execute(rw, data)
}
