package Repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"
)

type CategoryRepository interface {
	CreateCategory(Category *Model.Category) (*Model.Category, error)
	UpdateCategory(id int, Category *Model.Category) (*Model.Category, error)
	DeleteCategory(id int) error
	GetAllCategories() ([]*Model.Category, error)
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (c *categoryRepository) CreateCategory(category *Model.Category) (*Model.Category, error) {
	query := `INSERT INTO categories (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := c.db.QueryRow(query, category.Name, category.Description, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return nil, err
	}
	category.ID = int(id)
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	return category, nil
}

func (c *categoryRepository) GetAllCategories() ([]*Model.Category, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM categories`
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []*Model.Category{}
	for rows.Next() {
		category := Model.Category{}
		if err := rows.Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (c *categoryRepository) UpdateCategory(id int, category *Model.Category) (*Model.Category, error) {
	query := `UPDATE categories SET name = $1, description = $2, updated_at = $3 WHERE id = $4 RETURNING *`

	var updatedCategory Model.Category
	row := c.db.QueryRow(query, category.Name, category.Description, time.Now(), id)

	err := row.Scan(&updatedCategory.ID, &updatedCategory.Name, &updatedCategory.Description, &updatedCategory.CreatedAt, &updatedCategory.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("categoria no encontrada para actualizar: %w", err)
		}
		return nil, fmt.Errorf("error al escanear la categoría actualizada: %w", err)
	}

	return &updatedCategory, nil
}

func (c *categoryRepository) DeleteCategory(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := c.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
