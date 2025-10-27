package Auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerarToken(email string, roleID int) (string, error) {
	secret := os.Getenv("API_SECRET")
	if secret == "" {
		return "", errors.New("API_SECRET no está configurada en el entorno")
	}

	claims := jwt.MapClaims{}
	claims["id"] = email
	claims["role_id"] = roleID
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("fallo al firmar el token JWT: %w", err)
	}
	return tokenString, nil
}

func ValidarToken(r *http.Request) (jwt.MapClaims, error) {
	secret := os.Getenv("API_SECRET")
	if secret == "" {
		return nil, errors.New("API_SECRET no está configurada en el entorno")
	}

	jwtToken, err := ExtraerToken(r)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, fmt.Errorf("metodo inesperado: %s", token.Header["alg"])
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token inválido o expirado")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("formato de claims inválido")
	}

	return claims, nil
}

func ExtraerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
			return parts[1], nil
		}
	}

	queryToken := r.URL.Query().Get("token")
	if queryToken != "" {
		return queryToken, nil
	}

	return "", errors.New("token de autenticación no encontrado")
}
