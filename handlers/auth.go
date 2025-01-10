package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tripatra-test-go/models"
	"tripatra-test-go/utils"
)

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		panic("Error loading .env file")
// 	}
// }

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if "12345" != user.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate token
	token, err := utils.GenerateToken("aziz@mail.com")
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
}
