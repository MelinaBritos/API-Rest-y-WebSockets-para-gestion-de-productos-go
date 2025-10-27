package Middleware

import (
	"context"
	"net/http"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Auth"
	"github.com/golang-jwt/jwt"
)

const RoleAdminID = 1
const ClaimsContextKey = "claims"

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		claims, err := Auth.ValidarToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)

		next(w, r.WithContext(ctx))
	}
}

func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		claims, ok := r.Context().Value(ClaimsContextKey).(jwt.MapClaims)
		if !ok {
			http.Error(w, "Error de contexto. Token no procesado.", http.StatusUnauthorized)
			return
		}

		userRoleID, ok := claims["role_id"].(float64)
		if !ok {
			http.Error(w, "Rol faltante en el token.", http.StatusForbidden)
			return
		}

		if int(userRoleID) != RoleAdminID {
			http.Error(w, "Acceso denegado. Se requiere el rol de Administrador.", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}
