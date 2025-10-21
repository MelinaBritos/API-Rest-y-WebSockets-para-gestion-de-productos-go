package service

import (
	"errors"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"
	repository "github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) ObtenerProductos() ([]*Model.Product, error) {

	return s.repo.GetAllProducts()
}

func (s *ProductService) ObtenerProductoPorID(id int) (*Model.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) CrearProducto(product Model.Product) (*Model.Product, error) {
	//validacion al crear producto
	if err := validarProducto(product); err != nil {
		return nil, err
	}

	return s.repo.CreateProduct(&product)
}

func (s *ProductService) ActualizarProducto(id int, product Model.Product) (*Model.Product, error) {
	//validacion al actualizar producto
	if err := validarProducto(product); err != nil {
		return nil, err
	}

	return s.repo.UpdateProduct(id, &product)
}

func (s *ProductService) EliminarProducto(id int) error {
	return s.repo.DeleteProduct(id)
}

func validarProducto(product Model.Product) error {
	if product.Name == "" {
		return errors.New("el nombre del producto no puede estar vacío")
	}
	if len(product.Name) < 3 || len(product.Name) > 100 {
		return errors.New("el nombre debe tener entre 3 y 100 caracteres")
	}
	if product.Price <= 0 {
		return errors.New("el precio del producto debe ser mayor que cero")
	}
	if len(product.Description) > 500 {
		return errors.New("la descripción no puede exceder los 500 caracteres")
	}
	if product.Stock < 0 {
		return errors.New("el stock no puede ser negativo")
	}
	return nil
}
