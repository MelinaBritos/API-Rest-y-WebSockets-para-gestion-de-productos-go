package Handler

import (
	"encoding/json"
	"net/http"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Service"
)

type SearchHandler struct {
	service *Service.SearchService
}

func NewSearchHandler(service *Service.SearchService) *SearchHandler {
	return &SearchHandler{service: service}
}

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	searchType := query.Get("type")

	if searchType == "" {
		http.Error(w, "Debe ingresar el tipo de búsqueda.", http.StatusBadRequest)
		return
	}

	filters := map[string]string{
		"search":    query.Get("q"),         // Término de búsqueda (ej: 'computadora')
		"page":      query.Get("page"),      // Paginación (ej: '1')
		"limit":     query.Get("limit"),     // Paginación (ej: '10')
		"sort_by":   query.Get("sort_by"),   // Ordenamiento (ej: 'price')
		"order":     query.Get("order"),     // Ordenamiento (ej: 'desc')
		"min_price": query.Get("min_price"), // Filtro específico para productos
		"max_price": query.Get("max_price"), // Filtro específico para productos
	}

	switch searchType {
	case "product":
		result, err := h.service.SearchProducts(filters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	case "category":
		result, err := h.service.SearchCategories(filters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	default:
		http.Error(w, "Tipo de búsqueda no válida", http.StatusBadRequest)
		return
	}

}
