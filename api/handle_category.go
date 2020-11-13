package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	model "../model"
)

func CategoryListHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/api/v1/categories" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		HandleCategoryList_GET(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func HandleCategoryList_GET(w http.ResponseWriter, r *http.Request) {
	categoryList, err := model.GlobalCategoryStore.GetCategoryList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err2 := json.Marshal(categoryList)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func AllPostsByCategoryHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	if req.URL.Path != "/api/v1/posts-by-category" {
		http.Error(w, "Go back to the main page", 404)
		return
	}
	switch req.Method {
	case "GET":
		AllPostsByCategoryHandler_GET(w, req)
	default:
		http.Error(w, "Only GET method allowed, return to main page", 405)
	}
}

func AllPostsByCategoryHandler_GET(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query()["limit"]
	offset := r.URL.Query()["offset"]
	category := r.URL.Query()["category"]

	if limit == nil || len(limit) != 1 || offset == nil || len(offset) != 1 {
		// todo render 404
		http.Error(w, "Params required", 404)
		return
	}

	limitInt, err := strconv.Atoi(limit[0])
	if err != nil {
		http.Error(w, "Invalid limit param", 404)
		return
	}
	offsetInt, err := strconv.Atoi(offset[0])
	if err != nil {
		http.Error(w, "Id required", 404)
		return
	}
	categoryId, err := strconv.Atoi(category[0])
	if err != nil {
		http.Error(w, "Id required", 404)
		return
	}

	posts, err := model.GlobalPostStore.GetByCategory(offsetInt, limitInt, categoryId)
	if err != nil || posts == nil {
		http.Error(w, "Error", 404)
		return
	}

	js, err2 := json.Marshal(posts)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
