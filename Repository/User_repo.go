package Repository

import (
	"database/sql"
	"fmt"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"
)

type UserRepository interface {
	GetUserByEmail(email string) (*Model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetUserByEmail(email string) (*Model.User, error) {
	query := `SELECT id, username, email, password, role_id FROM Users WHERE email = $1`
	row := u.db.QueryRow(query, email)

	usuario := Model.User{}
	err := row.Scan(&usuario.ID, &usuario.Username, &usuario.Email, &usuario.Password, &usuario.RoleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario no encontrado: %w", err)
		}
		return nil, fmt.Errorf("error al escanear el usuario: %w", err)
	}

	return &usuario, nil
}
