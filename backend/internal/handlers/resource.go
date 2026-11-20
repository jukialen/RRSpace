package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"rrspace/backend/internal/middleware"
	"rrspace/backend/internal/models"
	"gorm.io/gorm"
)

type ResourceHandler struct {
	DB *gorm.DB
}

func NewResourceHandler(db *gorm.DB) *ResourceHandler {
	return &ResourceHandler{DB: db}
}

// ListResources — GET /studios/:id/resources
func (h *ResourceHandler) ListResources(w http.ResponseWriter, r *http.Request) {
	studioID := mux.Vars(r)["id"]
	var resources []models.Resource
	h.DB.Where("studio_id = ?", studioID).Find(&resources)
	json.NewEncoder(w).Encode(resources)
}

// CreateResource — POST /studios/:id/resources
func (h *ResourceHandler) CreateResource(w http.ResponseWriter, r *http.Request) {
	studioID := mux.Vars(r)["id"]
	var res models.Resource
	json.NewDecoder(r.Body).Decode(&res)
	res.StudioID = studioID
	h.DB.Create(&res)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// UpdateResource — PATCH /resources/:id
func (h *ResourceHandler) UpdateResource(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var res models.Resource
	h.DB.First(&res, "id = ?", id)
	json.NewDecoder(r.Body).Decode(&res)
	h.DB.Save(&res)
	json.NewEncoder(w).Encode(res)
}

// DeleteResource — DELETE /resources/:id
func (h *ResourceHandler) DeleteResource(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	h.DB.Delete(&models.Resource{}, "id = ?", id)
	w.WriteHeader(http.StatusNoContent)
}

func (h *ResourceHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/studios/{id}/resources", h.ListResources).Methods(http.MethodGet)

	owner := func(h http.Handler) http.Handler {
		return middleware.RequireRole("owner")(h)
	}
	r.Handle("/studios/{id}/resources", owner(http.HandlerFunc(h.CreateResource))).Methods(http.MethodPost)
	r.Handle("/resources/{id}", owner(http.HandlerFunc(h.UpdateResource))).Methods(http.MethodPatch)
	r.Handle("/resources/{id}", owner(http.HandlerFunc(h.DeleteResource))).Methods(http.MethodDelete)
}
