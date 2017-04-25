package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/signature"
)

func main() {
	// set addr
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()

	// setup the providers
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		github.New("d051e93fe220dd811c66", "fa473ba84c4b96b72aad7c006a76f0871c256d00", "http://127.0.0.1:8080/auth/callback/github"),
	)

	r := newRoom()
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	// get the room going
	go r.run()
	// start the web server
	log.Println("Starting web server on", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
