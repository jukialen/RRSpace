package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/studiobooking/go-api/internal/middleware"
	"github.com/studiobooking/go-api/internal/models"
	"gorm.io/gorm"
)

type BookingHandler struct {
	DB *gorm.DB
}

func NewBookingHandler(db *gorm.DB) *BookingHandler {
	return &BookingHandler{DB: db}
}

// CreateBooking — POST /bookings
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	json.NewDecoder(r.Body).Decode(&booking)

	// Walidacja: sprawdź czy slot jest wolny
	// (query do bookings z overlapping datami na ten sam resource)
	var conflict int64
	h.DB.Model(&models.Booking{}).
		Where("resource_id = ? AND status NOT IN ? AND start_at < ? AND end_at > ?",
			booking.ResourceID, []string{"cancelled"}, booking.EndAt, booking.StartAt).
		Count(&conflict)

	if conflict > 0 {
		http.Error(w, `{"error":"slot not available"}`, http.StatusConflict)
		return
	}

	booking.ClientID = middleware.GetUserFromContext(r).ID
	booking.Status = "pending"
	h.DB.Create(&booking)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

// ConfirmBooking — PATCH /bookings/:id/confirm
func (h *BookingHandler) ConfirmBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	h.DB.Model(&models.Booking{}).Where("id = ?", id).Update("status", "confirmed")
	w.WriteHeader(http.StatusOK)
}

// CancelBooking — PATCH /bookings/:id/cancel
func (h *BookingHandler) CancelBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	h.DB.Model(&models.Booking{}).Where("id = ?", id).Update("status", "cancelled")
	w.WriteHeader(http.StatusOK)
}

func (h *BookingHandler) RegisterRoutes(r *mux.Router) {
	auth := func(h http.Handler) http.Handler {
		return middleware.AuthMiddleware(h)
	}

	r.Handle("/bookings", auth(http.HandlerFunc(h.CreateBooking))).Methods(http.MethodPost)
	r.Handle("/bookings/{id}/confirm", auth(http.HandlerFunc(h.ConfirmBooking))).Methods(http.MethodPatch)
	r.Handle("/bookings/{id}/cancel", auth(http.HandlerFunc(h.CancelBooking))).Methods(http.MethodPatch)
}