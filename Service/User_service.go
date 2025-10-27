package Service

import (
	"fmt"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Auth"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo Repository.UserRepository
}

func NewUserService(repo Repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ValidarUsuario(credenciales Model.Credenciales) (string, error) {
	usuario, err := s.repo.GetUserByEmail(credenciales.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(credenciales.Password)); err != nil {
		return "", fmt.Errorf("contraseña incorrecta: %w", err)
	}

	token, err := Auth.GenerarToken(usuario.Email, usuario.RoleID)
	if err != nil {
		return "", fmt.Errorf("error al generar el token: %w", err)
	}
	return token, nil

}
