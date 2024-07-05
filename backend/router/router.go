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
	mux.Get("/api/categories/total", handler.Ctrl.GetTotalCategoryNumber)
	mux.Get("/api/categories/page/{size}", handler.Ctrl.GetTotalCategoryPageNumber)
	mux.Get("/api/categories/page/{number}/size/{size}", handler.Ctrl.GetCategoryByPage)
	mux.Post("/api/category", handler.Ctrl.AddCategory)
	mux.Get("/api/category/{id}", handler.Ctrl.GetCategoryById)
	mux.Put("/api/category/{id}", handler.Ctrl.UpdateCategory)
	mux.Delete("/api/category/{id}", handler.Ctrl.DeleteCategory)

	mux.Get("/api/category/{id}/subcategories", handler.Ctrl.GetAllSubCategory)
	mux.Get("/api/category/{id}/subcategories/total", handler.Ctrl.GetTotalSubCategoryNumber)
	mux.Get("/api/category/{id}/subcategories/page/{size}", handler.Ctrl.GetTotalSubCategoryPageNumber)
	mux.Get("/api/category/{id}/subcategories/page/{number}/size/{size}", handler.Ctrl.GetSubCategoryByPage)
	mux.Post("/api/category/{id}/subcategory", handler.Ctrl.AddSubCategory)
	mux.Get("/api/category/{id}/subcategory/{subId}", handler.Ctrl.GetSubCategoryById)
	mux.Put("/api/category/{id}/subcategory/{subId}", handler.Ctrl.UpdateSubCategory)
	mux.Delete("/api/category/{id}/subcategory/{subId}", handler.Ctrl.DeleteSubCategory)

	// Color routes
	mux.Get("/api/color-categories", handler.Ctrl.GetAllColorCategory)
	mux.Get("/api/colors", handler.Ctrl.GetAllColor)

	// Tag routes
	mux.Get("/api/tags", handler.Ctrl.GetAllTag)
	mux.Get("/api/tags/total", handler.Ctrl.GetTotalTagNumber)
	mux.Get("/api/tags/page/{size}", handler.Ctrl.GetTotalTagPageNumber)
	mux.Get("/api/tags/page/{number}/size/{size}", handler.Ctrl.GetTagsByPage)
	mux.Post("/api/tag", handler.Ctrl.AddTag)
	mux.Get("/api/tag/{id}", handler.Ctrl.GetTagById)
	mux.Put("/api/tag/{id}", handler.Ctrl.UpdateTag)
	mux.Delete("/api/tag/{id}", handler.Ctrl.DeleteTag)

	// Record
	mux.Post("/api/record/upload-image", handler.Ctrl.UploadRecordImage)

	// File server
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return cors(mux)
}