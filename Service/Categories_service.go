package Service

import (
	"encoding/json"
	"errors"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Repository"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/WebSocket"
)

type CategoryService struct {
	repo Repository.CategoryRepository
}

func NewCategoryService(repo Repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) ObtenerCategorias() ([]*Model.Category, error) {

	return s.repo.GetAllCategories()
}

func (s *CategoryService) CrearCategoria(category Model.Category) (*Model.Category, error) {
	//validacion al crear categoria
	if err := validarCategoria(category); err != nil {
		return nil, err
	}

	newCategory, err := s.repo.CreateCategory(&category)
	if err != nil {
		return nil, err
	}

	// Emitir evento de creación de categoría a través del WebSocket
	event := map[string]interface{}{
		"data":   newCategory,
		"event":  "CATEGORY_CREATED",
		"action": "Se creó una nueva categoría",
	}
	eventJSON, _ := json.Marshal(event)
	WebSocket.Emit(eventJSON)

	return newCategory, nil
}

func (s *CategoryService) ActualizarCategoria(id int, category Model.Category) (*Model.Category, error) {
	//validacion al actualizar categoria
	if err := validarCategoria(category); err != nil {
		return nil, err
	}

	updatedCategory, err := s.repo.UpdateCategory(id, &category)
	if err != nil {
		return nil, err
	}

	event := map[string]interface{}{
		"data":   updatedCategory,
		"event":  "CATEGORY_UPDATED",
		"action": "Se actualizó una categoría",
	}
	eventJSON, _ := json.Marshal(event)
	WebSocket.Emit(eventJSON)

	return updatedCategory, nil
}

func (s *CategoryService) EliminarCategoria(id int) error {
	event := map[string]interface{}{
		"event":  "CATEGORY_DELETED",
		"action": "Se eliminó una categoría",
		"id":     id,
	}
	eventJSON, _ := json.Marshal(event)
	WebSocket.Emit(eventJSON)

	return s.repo.DeleteCategory(id)
}

func validarCategoria(category Model.Category) error {
	if category.Name == "" {
		return errors.New("el nombre de la categoría no puede estar vacío")
	}
	if len(category.Name) < 3 || len(category.Name) > 100 {
		return errors.New("el nombre debe tener entre 3 y 100 caracteres")
	}
	if len(category.Description) > 500 {
		return errors.New("la descripción no puede exceder los 500 caracteres")
	}
	return nil
}
