package Handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Service"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	service *Service.ProductService
}

func NewProductHandler(service *Service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (p *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := p.service.ObtenerProductos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir productos a JSON y escribir en la respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (p *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	product, err := p.service.ObtenerProductoPorID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Convertir producto a JSON y escribir en la respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (p *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := p.service.CrearProducto(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}

func (p *ProductHandler) CreateProducts(w http.ResponseWriter, r *http.Request) {
	var products []Model.Product
	if err := json.NewDecoder(r.Body).Decode(&products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdProducts, err := p.service.CrearProductos(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdProducts)
}

func (p *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	var product Model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := p.service.ActualizarProducto(id, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (p *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	if err := p.service.EliminarProducto(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (p *ProductHandler) GetProductHistory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	queryParams := r.URL.Query()
	startDate := queryParams.Get("start")
	endDate := queryParams.Get("end")

	history, err := p.service.ObtenerHistorialProducto(id, startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}
