package router

import (
	"database/sql"
	"encoding/json"
	"github.com/margar-melkonyan/watch-later.git/internal/repository"
	"github.com/margar-melkonyan/watch-later.git/internal/service/categories"
	"net/http"
	"strconv"
	"strings"
)

var categoryService *service.CategoryService

type Response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func categoriesRoutes(db *sql.DB) *http.ServeMux {
	categories := http.NewServeMux()
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService = service.NewCategoryService(categoryRepo)

	categories.HandleFunc("GET /", getCategories)
	categories.HandleFunc("GET /{id}", getCategoryById)
	categories.HandleFunc("POST /", storeCategory)
	categories.HandleFunc("POST /{id}/restore", func(w http.ResponseWriter, r *http.Request) {

	})
	categories.HandleFunc("PUT /{id}", func(w http.ResponseWriter, r *http.Request) {

	})
	categories.HandleFunc("DELETE /{id}", func(w http.ResponseWriter, r *http.Request) {

	})

	return categories
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := categoryService.GetCategories()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data := Response{
		Data: categories,
	}

	response, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
}

func getCategoryById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(strings.Split(r.URL.Path, "/")[1], 0, 64)
	if err != nil {
		errResponse := ErrorResponse{
			Message: "Invalid category `id`",
		}

		response, err := json.Marshal(errResponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(response))
		return
	}

	category, err := categoryService.GetCategory(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data := Response{
		Data: category,
	}

	response, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
}

func storeCategory(w http.ResponseWriter, r *http.Request) {

}
