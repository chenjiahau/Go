package router

import (
	"net/http"

	"example.com/project/page"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetRouter() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(WriteToConsole)

	mux.Get("/", page.Index)

	return mux
}
