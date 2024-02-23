package main

import (
	"github/janislaus/wogalo/pkg/config"
	"github/janislaus/wogalo/pkg/handlers"
	"github/janislaus/wogalo/pkg/render"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const port = ":8080"

func main() {
	// var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache.")
	}

	config := config.AppConfig{TemplateCache: tc, UseCache: false, Store: sessions.NewCookieStore([]byte("your-secret-key"))}
	repo := handlers.Repository{Config: &config}

	serve := &http.Server{
		Addr:    port,
		Handler: routes(&repo),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)
}
