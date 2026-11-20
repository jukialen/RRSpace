package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"rrspace/backend/internal/middleware"
	"rrspace/backend/internal/models"
	"gorm.io/gorm"
)

type StudioHandler struct {
	DB *gorm.DB
}

func NewStudioHandler(db *gorm.DB) *StudioHandler {
	return &StudioHandler{DB: db}
}

// CreateStudio — POST /studios
func (h *StudioHandler) CreateStudio(w http.ResponseWriter, r *http.Request) {
	var studio models.Studio
	if err := json.NewDecoder(r.Body).Decode(&studio); err != nil {
		http.Error(w, `{"error":"invalid body"}`, http.StatusBadRequest)
		return
	}
	studio.OwnerID = middleware.GetUserFromContext(r).ID
	studio.IsActive = true

	if err := h.DB.Create(&studio).Error; err != nil {
		http.Error(w, `{"error":`+strconv.Quote(err.Error())+`}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(studio)
}

// GetStudio — GET /studios/:id
func (h *StudioHandler) GetStudio(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var studio models.Studio
	if err := h.DB.Preload("Resources").Preload("Photos").First(&studio, "id = ?", id).Error; err != nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(studio)
}

// UpdateStudio — PATCH /studios/:id
func (h *StudioHandler) UpdateStudio(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var studio models.Studio
	if err := h.DB.First(&studio, "id = ?", id).Error; err != nil {
		http.NotFound(w, r)
		return
	}
	json.NewDecoder(r.Body).Decode(&studio)
	h.DB.Save(&studio)
	json.NewEncoder(w).Encode(studio)
}

// DeleteStudio — DELETE /studios/:id
func (h *StudioHandler) DeleteStudio(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	h.DB.Delete(&models.Studio{}, "id = ?", id)
	w.WriteHeader(http.StatusNoContent)
}

// UploadPhotos — POST /studios/:id/photos (multipart)
func (h *StudioHandler) UploadPhotos(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	// Walidacja: max 5MB, typy: image/jpeg, image/png
	file, _, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, `{"error":"file required"}`, http.StatusBadRequest)
		return
	}
	defer file.Close()

	// TODO: upload do R2
	_ = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"url": "placeholder-url"})
}

func (h *StudioHandler) RegisterRoutes(r *mux.Router) {
	// Publiczne
	r.HandleFunc("/studios", h.CreateStudio).Methods(http.MethodPost)
	r.HandleFunc("/studios/{id}", h.GetStudio).Methods(http.MethodGet)

	// Właściciel
	owner := func(h http.Handler) http.Handler {
		return middleware.RequireRole("owner")(h)
	}
	r.Handle("/studios/{id}", owner(http.HandlerFunc(h.UpdateStudio))).Methods(http.MethodPatch)
	r.Handle("/studios/{id}", owner(http.HandlerFunc(h.DeleteStudio))).Methods(http.MethodDelete)
	r.Handle("/studios/{id}/photos", owner(http.HandlerFunc(h.UploadPhotos))).Methods(http.MethodPost)
}