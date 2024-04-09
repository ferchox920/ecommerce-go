package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/fercho920/ecommerce-go/constants"
	"github.com/fercho920/ecommerce-go/db"
	"github.com/fercho920/ecommerce-go/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	result := db.DB.First(&user, params["id"])
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	if err := db.DB.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&users); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userDto models.UserCreate

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
	
		http.Error(w, "Invalid request body format", http.StatusBadRequest)
		return
	}

	existingUser := models.User{}
	result := db.DB.Where("email = ?", userDto.Email).First(&existingUser)
	if result.Error != nil && result.RowsAffected != 0 {
		http.Error(w, "Error checking for existing user", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected != 0 {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}


	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error encrypting password", http.StatusInternalServerError)
		return
	}

	// Create a new user
	newUser := models.User{
		ID:        uuid.New(),
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Password:  string(hashedPassword),
		Email:     userDto.Email,
		Phone:     userDto.Phone,
	}

	// Save the new user
	result = db.DB.Create(&newUser)
	if result.Error != nil {
		// Handle specific database errors (e.g., constraint violations)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Omit password field in the response
	newUser.Password = ""

	// Success response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&newUser); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(constants.UserContextKey).(models.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}
	var updateUser models.UserUpdateClient

	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if updateUser.FirstName != "" {
		user.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		user.LastName = updateUser.LastName
	}
	if updateUser.Phone != "" {
		user.Phone = updateUser.Phone
	}

	if err := db.DB.Save(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := db.DB.Delete(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

}
