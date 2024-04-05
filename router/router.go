package router

import (
	"net/http"

	"example.com/project/page"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetRoutes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(NoSurf)
	mux.Use(WriteToConsole)

	// Static files
	staticServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", staticServer))

	mux.Get("/", page.Repo.Index)
	mux.Get("/message", page.Repo.Message)
	mux.Post("/send-message", page.Repo.SendMessage)

	mux.Post("/login", page.Repo.Login)
	mux.Get("/login/users", page.Repo.GetUsers)

	return mux
}

