package main

import (
	"github/janislaus/wogalo/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(repo *handlers.Repository) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad(repo.Config.Store))
	mux.Get("/", repo.Home)
	mux.Get("/about", repo.About)

	return mux
}
