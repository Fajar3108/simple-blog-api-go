package auth

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *Handler {
	return &Handler{authService: authService}
}

func (ah Handler) Signin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Login user...")
}
