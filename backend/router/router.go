package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"ivanfun.com/mis/handler"
)

func GetRoutes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(EnableCORS)
	mux.Use(WriteToConsole)
	mux.Use(ParseAuthorization)

	mux.Get("/api", handler.Ctrl.Index)
	mux.Post("/api/sign-up", handler.Ctrl.SignUp)
	mux.Post("/api/sign-in", handler.Ctrl.SignIn)

	// Protected routes
	mux.Get("/api/auth/test-protecting-route", handler.Ctrl.TestProtectingRoute)
	mux.Get("/api/auth/sign-out", handler.Ctrl.SignOut)

	return mux
}