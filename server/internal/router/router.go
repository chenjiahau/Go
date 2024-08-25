package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"ivanfun.com/mis/internal/handler"
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
	mux.Route("/api/auth", func(mux chi.Router) {
		mux.Use(CheckTokenAlive)

		mux.Get("/verify-token", handler.Ctrl.VerifyToken)
		mux.Get("/sign-out", handler.Ctrl.SignOut)

		// Member routes
		mux.Get("/members", handler.Ctrl.GetAllMember)
		mux.Get("/member-roles", handler.Ctrl.GetAllMemberRole)
		mux.Get("/members", handler.Ctrl.GetAllMember)
		mux.Get("/members/total", handler.Ctrl.GetTotalMemberNumber)
		mux.Get("/members/page/{size}", handler.Ctrl.GetTotalMemberPageNumber)
		mux.Get("/members/page/{number}/size/{size}", handler.Ctrl.GetMemberByPage)
		mux.Post("/member", handler.Ctrl.AddMember)
		mux.Get("/member/{id}", handler.Ctrl.GetMemberById)
		mux.Put("/member/{id}", handler.Ctrl.UpdateMember)
		mux.Delete("/member/{id}", handler.Ctrl.DeleteMember)

		// Category routes
		mux.Get("/categories", handler.Ctrl.GetAllCategory)
		mux.Get("/categories/total", handler.Ctrl.GetTotalCategoryNumber)
		mux.Get("/categories/page/{size}", handler.Ctrl.GetTotalCategoryPageNumber)
		mux.Get("/categories/page/{number}/size/{size}", handler.Ctrl.GetCategoryByPage)
		mux.Post("/category", handler.Ctrl.AddCategory)
		mux.Get("/category/{id}", handler.Ctrl.GetCategoryById)
		mux.Put("/category/{id}", handler.Ctrl.UpdateCategory)
		mux.Delete("/category/{id}", handler.Ctrl.DeleteCategory)
		mux.Get("/category/{id}/subcategories", handler.Ctrl.GetAllSubCategory)
		mux.Get("/category/{id}/subcategories/total", handler.Ctrl.GetTotalSubCategoryNumber)
		mux.Get("/category/{id}/subcategories/page/{size}", handler.Ctrl.GetTotalSubCategoryPageNumber)
		mux.Get("/category/{id}/subcategories/page/{number}/size/{size}", handler.Ctrl.GetSubCategoryByPage)
		mux.Post("/category/{id}/subcategory", handler.Ctrl.AddSubCategory)
		mux.Get("/category/{id}/subcategory/{subId}", handler.Ctrl.GetSubCategoryById)
		mux.Put("/category/{id}/subcategory/{subId}", handler.Ctrl.UpdateSubCategory)
		mux.Delete("/category/{id}/subcategory/{subId}", handler.Ctrl.DeleteSubCategory)

		// Color routes
		mux.Get("/color-categories", handler.Ctrl.GetAllColorCategory)
		mux.Get("/colors", handler.Ctrl.GetAllColor)

		// Tag routes
		mux.Get("/tags", handler.Ctrl.GetAllTag)
		mux.Get("/tags/total", handler.Ctrl.GetTotalTagNumber)
		mux.Get("/tags/page/{size}", handler.Ctrl.GetTotalTagPageNumber)
		mux.Get("/tags/page/{number}/size/{size}", handler.Ctrl.GetTagsByPage)
		mux.Post("/tag", handler.Ctrl.AddTag)
		mux.Get("/tag/{id}", handler.Ctrl.GetTagById)
		mux.Put("/tag/{id}", handler.Ctrl.UpdateTag)
		mux.Delete("/tag/{id}", handler.Ctrl.DeleteTag)

		// Record
		mux.Post("/record/upload-image", handler.Ctrl.UploadRecordImage)

		// Document routes
		mux.Get("/documents", handler.Ctrl.GetAllDocument)
		mux.Get("/documents/total", handler.Ctrl.GetTotalDocumentNumber)
		mux.Get("/documents/page/{size}", handler.Ctrl.GetTotalDocumentPageNumber)
		mux.Get("/documents/page/{number}/size/{size}", handler.Ctrl.GetDocumentByPage)
		mux.Post("/document", handler.Ctrl.AddDocument)
		mux.Get("/document/{id}", handler.Ctrl.GetDocumentById)
		mux.Put("/document/{id}", handler.Ctrl.UpdateDocument)
		mux.Delete("/document/{id}", handler.Ctrl.DeleteDocument)
		mux.Get("/document/{id}/comments", handler.Ctrl.GetAllDocumentComment)
		mux.Post("/document/{id}/comment", handler.Ctrl.AddDocumentComment)
		mux.Get("/document/{id}/comment/{commentId}", handler.Ctrl.GetDocumentCommentById)
		mux.Put("/document/{id}/comment/{commentId}", handler.Ctrl.UpdateDocumentComment)
		mux.Delete("/document/{id}/comment/{commentId}", handler.Ctrl.DeleteDocumentComment)
	})

	// Index route
	mux.Get("/", handler.Ctrl.Index)

	// File server
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/*", http.StripPrefix("/", fileServer))

	return mux
}