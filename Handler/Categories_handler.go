package Handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"
	service "github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Service"
	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (c *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.service.ObtenerCategorias()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir categorías a JSON y escribir en la respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (c *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category Model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := c.service.CrearCategoria(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}

func (c *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var category Model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := c.service.ActualizarCategoria(id, category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (c *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	if err := c.service.EliminarCategoria(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
