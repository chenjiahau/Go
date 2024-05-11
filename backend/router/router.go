package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/handlers"
	"ivanfun.com/mis/handler"
)

func GetRoutes() http.Handler {
	cors := handlers.CORS(
    handlers.AllowedHeaders([]string{"content-type"}),
    handlers.AllowedOrigins([]string{"*"}),
    handlers.AllowCredentials(),
	)

	mux := chi.NewRouter()
	mux.Use(WriteToConsole)
	mux.Use(ParseAuthorization)

	mux.Get("/api", handler.Ctrl.Index)
	mux.Post("/api/sign-up", handler.Ctrl.SignUp)
	mux.Post("/api/sign-in", handler.Ctrl.SignIn)

	// Protected routes
	mux.Get("/api/auth/verify-token", handler.Ctrl.VerifyToken)
	mux.Get("/api/auth/sign-out", handler.Ctrl.SignOut)

	return cors(mux)
}