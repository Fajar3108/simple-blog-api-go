package auth

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"simple-blog-api-golang/internal/auth/auth_requests"
	"simple-blog-api-golang/internal/auth/auth_resources"
	"simple-blog-api-golang/pkg/resource"
)

type Handler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *Handler {
	return &Handler{authService: authService}
}

func (ah *Handler) Signin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userRequest *auth_requests.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	token, err := ah.authService.Signin(r.Context(), userRequest)

	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(auth_resources.NewLoginResource(*resource.NewGeneralResource(200, "Login success"), token))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (ah *Handler) Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var registerRequest *auth_requests.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	token, err := ah.authService.Signup(r.Context(), registerRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(auth_resources.NewLoginResource(*resource.NewGeneralResource(201, "Register success"), token))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
