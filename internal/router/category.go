package router

import (
	"database/sql"
	"net/http"

	"github.com/margar-melkonyan/watch-later.git/internal/handler/controller"
)

func categoriesRoutes(db *sql.DB) *http.ServeMux {
	categories := http.NewServeMux()
	categoryController := controller.NewCategoryController(db)

	categories.HandleFunc("GET /", categoryController.GetCategories)
	categories.HandleFunc("GET /{id}", categoryController.GetCategoryById)
	categories.HandleFunc("POST /", categoryController.StoreCategory)
	categories.HandleFunc("POST /{id}/restore", categoryController.RestoreCategory)
	categories.HandleFunc("PUT /{id}", categoryController.UpdateCategory)
	categories.HandleFunc("DELETE /{id}", categoryController.DeleteCategory)

	return categories
}
