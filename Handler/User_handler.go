package Handler

import (
	"encoding/json"
	"net/http"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Service"
)

type UserHandler struct {
	service *Service.UserService
}

func NewUserHandler(service *Service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	var credenciales Model.Credenciales
	if err := json.NewDecoder(r.Body).Decode(&credenciales); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.service.ValidarUsuario(credenciales)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
