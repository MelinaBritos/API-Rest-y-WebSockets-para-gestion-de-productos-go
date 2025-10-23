package Service

import (
	"errors"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"
	repository "github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Repository"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
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

	return s.repo.CreateCategory(&category)
}

func (s *CategoryService) ActualizarCategoria(id int, category Model.Category) (*Model.Category, error) {
	//validacion al actualizar categoria
	if err := validarCategoria(category); err != nil {
		return nil, err
	}

	return s.repo.UpdateCategory(id, &category)
}

func (s *CategoryService) EliminarCategoria(id int) error {
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
