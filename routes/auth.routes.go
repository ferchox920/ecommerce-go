package routes

import (
	"encoding/json"
	"net/http"

	"github.com/fercho920/ecommerce-go/db"
	encryption "github.com/fercho920/ecommerce-go/encription"
	"github.com/fercho920/ecommerce-go/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := encryption.SignedLoginToken(&user)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"AccessToken": tokenString}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
