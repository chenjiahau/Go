package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/handlers"
	"ivanfun.com/mis/handler"
)

func GetRoutes() http.Handler {
	cors := handlers.CORS(
    handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
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

	// Category routes
	mux.Get("/api/categories", handler.Ctrl.GetAllCategory)
	mux.Post("/api/category", handler.Ctrl.AddCategory)
	mux.Get("/api/category/{id}", handler.Ctrl.GetCategoryById)
	mux.Put("/api/category/{id}", handler.Ctrl.UpdateCategory)
	mux.Delete("/api/category/{id}", handler.Ctrl.DeleteCategory)

	mux.Get("/api/category/{id}/subcategories", handler.Ctrl.GetAllSubCategory)
	mux.Post("/api/category/{id}/subcategory", handler.Ctrl.AddSubCategory)
	mux.Get("/api/category/{id}/subcategory/{subId}", handler.Ctrl.GetSubCategoryById)
	mux.Put("/api/category/{id}/subcategory/{subId}", handler.Ctrl.UpdateSubCategory)
	mux.Delete("/api/category/{id}/subcategory/{subId}", handler.Ctrl.DeleteSubCategory)

	return cors(mux)
}