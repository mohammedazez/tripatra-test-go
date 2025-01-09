package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"tripatra-test-go/utils"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Println("email", email)
	fmt.Println("os email", os.Getenv("email"))
	if email == os.Getenv("email") && password == os.Getenv("password") {
		token, _ := utils.GenerateJWT(email)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]string{"token": token}

		json.NewEncoder(w).Encode(response)
		return
	}
	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}
