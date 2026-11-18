package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/studiobooking/go-api/internal/middleware"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler { return &AuthHandler{} }

// Register — rejestracja (deleguje do Appwrite)
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: wywołanie Appwrite SDK: CreateUser
	// Następnie zapis w lokalnej tabeli users (role)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// Login — logowanie (Appwrite sessions)
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// Logout — invalidacja sesji
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// Me — GET /users/me
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUserFromContext(r)
	if user == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateMe — PATCH /users/me
func (h *AuthHandler) UpdateMe(w http.ResponseWriter, r *http.Request) {
	// TODO: walidacja, zapis do Appwrite + lokalnie
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// RegisterRoutes
func (h *AuthHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/auth/register", h.Register).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", h.Login).Methods(http.MethodPost)

	// Zabezpieczone
	r.Handle("/auth/logout", h.wrappedLogout()).Methods(http.MethodPost)
	r.Handle("/users/me", h.wrappedMe()).Methods(http.MethodGet)
	r.Handle("/users/me", h.wrappedUpdateMe()).Methods(http.MethodPatch)
}

func (h *AuthHandler) wrappedLogout() http.Handler {
	return middleware.AuthMiddleware(http.HandlerFunc(h.Logout))
}
func (h *AuthHandler) wrappedMe() http.Handler {
	return middleware.AuthMiddleware(http.HandlerFunc(h.Me))
}
func (h *AuthHandler) wrappedUpdateMe() http.Handler {
	return middleware.AuthMiddleware(http.HandlerFunc(h.UpdateMe))
}