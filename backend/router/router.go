package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"ivanfun.com/mis/handler"
)

func GetRoutes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(WriteToConsole)
	mux.Use(ParseAuthorization)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization", "X-Requested-With", "Accept"},
	}))

	mux.Get("/api", handler.Ctrl.Index)
	mux.Post("/api/sign-up", handler.Ctrl.SignUp)
	mux.Post("/api/sign-in", handler.Ctrl.SignIn)

	// Protected routes
	mux.Get("/api/auth/verify-token", handler.Ctrl.VerifyToken)
	mux.Get("/api/auth/sign-out", handler.Ctrl.SignOut)

	// Member routes
	mux.Get("/api/member-roles", handler.Ctrl.GetAllMemberRole)
	mux.Get("/api/members", handler.Ctrl.GetAllMember)
	mux.Get("/api/members/total", handler.Ctrl.GetTotalMemberNumber)
	mux.Get("/api/members/page/{size}", handler.Ctrl.GetTotalMemberPageNumber)
	mux.Get("/api/members/page/{number}/size/{size}", handler.Ctrl.GetMemberByPage)
	mux.Post("/api/member", handler.Ctrl.AddMember)
	mux.Get("/api/member/{id}", handler.Ctrl.GetMemberById)
	mux.Put("/api/member/{id}", handler.Ctrl.UpdateMember)
	mux.Delete("/api/member/{id}", handler.Ctrl.DeleteMember)

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

	// Document routes
	mux.Get("/api/documents", handler.Ctrl.GetAllDocument)
	mux.Get("/api/documents/total", handler.Ctrl.GetTotalDocumentNumber)
	mux.Get("/api/documents/page/{size}", handler.Ctrl.GetTotalDocumentPageNumber)
	mux.Get("/api/documents/page/{number}/size/{size}", handler.Ctrl.GetDocumentByPage)
	mux.Post("/api/document", handler.Ctrl.AddDocument)
	mux.Get("/api/document/{id}", handler.Ctrl.GetDocumentById)
	mux.Put("/api/document/{id}", handler.Ctrl.UpdateDocument)
	mux.Delete("/api/document/{id}", handler.Ctrl.DeleteDocument)

	mux.Get("/api/document/{id}/comments", handler.Ctrl.GetAllDocumentComment)
	mux.Post("/api/document/{id}/comment", handler.Ctrl.AddDocumentComment)
	mux.Get("/api/document/{id}/comment/{commentId}", handler.Ctrl.GetDocumentCommentById)
	mux.Put("/api/document/{id}/comment/{commentId}", handler.Ctrl.UpdateDocumentComment)
	mux.Delete("/api/document/{id}/comment/{commentId}", handler.Ctrl.DeleteDocumentComment)

	return mux
}