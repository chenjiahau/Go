package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"ivanfun.com/mis/handler"
)

func GetRoutes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(WriteToConsole)

	mux.Get("/api", handler.Ctrl.Index)

	return mux
}