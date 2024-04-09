package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/fercho920/ecommerce-go/constants"
	"github.com/fercho920/ecommerce-go/db"
	"github.com/fercho920/ecommerce-go/models"
	"github.com/golang-jwt/jwt/v5"
)

// UserContextKey es una clave para almacenar el usuario en el contexto del request.


func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Token de autorización faltante", http.StatusUnauthorized)
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Formato de token de autorización inválido", http.StatusUnauthorized)
            return
        }
        tokenString := parts[1]

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("secret"), nil // Reemplazar "secret" con tu clave secreta real
        })

        if err != nil || !token.Valid {
            http.Error(w, "Token de autorización inválido", http.StatusUnauthorized)
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            http.Error(w, "Formato de reclamaciones de token inválido", http.StatusUnauthorized)
            return
        }

        email, emailExists := claims["email"].(string)
        if !emailExists {
            http.Error(w, "Falta reclamación de correo electrónico en el token", http.StatusUnauthorized)
            return
        }

        var user models.User
        if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
            http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
            return
        }

	
        ctx := context.WithValue(r.Context(), constants.UserContextKey, user)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
