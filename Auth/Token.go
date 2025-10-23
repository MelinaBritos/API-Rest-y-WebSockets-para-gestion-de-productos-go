package Auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerarToken(userID int, roleID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = userID
	claims["role_id"] = roleID
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func ValidarToken(r *http.Request) error {
	jwtToken := ExtraerToken(r)
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(os.Getenv("API_SECRET")), nil
		}
		return nil, fmt.Errorf("metodo inesperado: %s", token.Header["alg"])
	})

	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}

	return nil

}

func Pretty(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error al formatear JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func ExtraerToken(r *http.Request) string {
	parametros := r.URL.Query()
	token := parametros.Get("token")
	if token != "" {
		return token
	}

	tokenString := r.Header.Get("Authorization")
	if len(strings.Split(tokenString, " ")) == 2 {
		return strings.Split(tokenString, " ")[1]
	}
	return ""
}
