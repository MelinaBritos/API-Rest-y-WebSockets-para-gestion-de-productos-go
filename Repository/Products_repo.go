package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"
)

type ProductRepository interface {
	CreateProduct(product *Model.Product) (*Model.Product, error)
	GetProductByID(id int) (*Model.Product, error)
	UpdateProduct(id int, product *Model.Product) (*Model.Product, error)
	DeleteProduct(id int) error
	GetAllProducts() ([]*Model.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (p *productRepository) CreateProduct(product *Model.Product) (*Model.Product, error) {
	query := `INSERT INTO products (name, description, price, stock, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var id int
	err := p.db.QueryRow(query, product.Name, product.Description, product.Price, product.Stock, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return nil, err
	}
	product.ID = int(id)
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	return product, nil
}

func (p *productRepository) GetProductByID(id int) (*Model.Product, error) {
	query := `SELECT id, name, description, price, stock, created_at, updated_at FROM products WHERE id = $1`
	row := p.db.QueryRow(query, id)

	product := Model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) GetAllProducts() ([]*Model.Product, error) {
	query := `SELECT id, name, description, price, stock, created_at, updated_at FROM products`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*Model.Product{}
	for rows.Next() {
		product := Model.Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (p *productRepository) UpdateProduct(id int, product *Model.Product) (*Model.Product, error) {
	query := `UPDATE products SET name = $1, description = $2, price = $3, stock = $4, updated_at = $5 WHERE id = $6`

	result, err := p.db.Exec(query, product.Name, product.Description, product.Price, product.Stock, time.Now(), id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("product not found for update")
	}
	product.ID = id
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	return product, nil
}

func (p *productRepository) DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
