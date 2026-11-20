package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"rrspace/backend/internal/middleware"
	"rrspace/backend/internal/models"
	"gorm.io/gorm"
)

type ScheduleHandler struct {
	DB *gorm.DB
}

func NewScheduleHandler(db *gorm.DB) *ScheduleHandler {
	return &ScheduleHandler{DB: db}
}

// ListSchedules — GET /studios/:id/resources/:rid/schedules
func (h *ScheduleHandler) ListSchedules(w http.ResponseWriter, r *http.Request) {
	rid := mux.Vars(r)["rid"]
	var schedules []models.Schedule
	h.DB.Where("resource_id = ?", rid).Find(&schedules)
	json.NewEncoder(w).Encode(schedules)
}

// CreateSchedule — POST /studios/:id/resources/:rid/schedules
func (h *ScheduleHandler) CreateSchedule(w http.ResponseWriter, r *http.Request) {
	rid := mux.Vars(r)["rid"]
	var sched models.Schedule
	json.NewDecoder(r.Body).Decode(&sched)
	sched.ResourceID = rid
	h.DB.Create(&sched)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sched)
}

// GenerateSlots — GET /studios/:id/resources/:rid/slots?date=2026-11-15
// Algorytm generowania slotów z uwzględnieniem exceptionów
func (h *ScheduleHandler) GenerateSlots(w http.ResponseWriter, r *http.Request) {
	rid := mux.Vars(r)["rid"]
	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		http.Error(w, `{"error":"date param required"}`, http.StatusBadRequest)
		return
	}
	day, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, `{"error":"invalid date"}`, http.StatusBadRequest)
		return
	}

	// 1. Pobierz harmonogram tygodniowy
	var schedules []models.Schedule
	h.DB.Where("resource_id = ? AND day_of_week = ?", rid, int(day.Weekday())).Find(&schedules)

	// 2. Pobierz exceptiony na ten dzień
	var exceptions []models.ScheduleException
	h.DB.Where("resource_id = ? AND date_from <= ? AND date_to >= ?", rid, day, day).Find(&exceptions)

	// 3. Generuj sloty
	slots := generateSlots(schedules, exceptions, day)

	json.NewEncoder(w).Encode(slots)
}

// AddException — POST /resources/:id/exceptions
func (h *ScheduleHandler) AddException(w http.ResponseWriter, r *http.Request) {
	rid := mux.Vars(r)["id"]
	var ex models.ScheduleException
	json.NewDecoder(r.Body).Decode(&ex)
	ex.ResourceID = rid
	h.DB.Create(&ex)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ex)
}

// generateSlots — rdzeń algorytmu
func generateSlots(schedules []models.Schedule, exceptions []models.ScheduleException, day time.Time) []models.TimeSlot {
	var slots []models.TimeSlot

	// Sprawdź czy jest blokada
	for _, ex := range exceptions {
		if ex.IsBlocked {
			return slots // puste — cały dzień zablokowany
		}
	}

	for _, sched := range schedules {
		openT, _ := time.Parse("15:04", sched.OpenTime)
		closeT, _ := time.Parse("15:04", sched.CloseTime)

		// Nadpisanie z exceptionów
		for _, ex := range exceptions {
			if ex.OverrideOpenTime != nil {
				openT, _ = time.Parse("15:04", *ex.OverrideOpenTime)
			}
			if ex.OverrideCloseTime != nil {
				closeT, _ = time.Parse("15:04", *ex.OverrideCloseTime)
			}
		}

		dur := time.Duration(sched.SlotDurationMinutes) * time.Minute
		for t := openT; t.Add(dur).Hour() <= closeT.Hour(); t = t.Add(dur) {
			start := time.Date(day.Year(), day.Month(), day.Day(), t.Hour(), t.Minute(), 0, 0, time.Local)
			end := start.Add(dur)

			slot := models.TimeSlot{
				ResourceID: sched.ResourceID,
				Date:       day.Format("2006-01-02"),
				StartAt:    start,
				EndAt:      end,
				Available:  true, // TODO: sprawdź collision z bookings
			}
			slots = append(slots, slot)
		}
	}
	return slots
}

func (h *ScheduleHandler) RegisterRoutes(r *mux.Router) {
	owner := func(h http.Handler) http.Handler {
		return middleware.RequireRole("owner")(h)
	}

	r.HandleFunc("/studios/{id}/resources/{rid}/schedules", h.ListSchedules).Methods(http.MethodGet)
	r.Handle("/studios/{id}/resources/{rid}/schedules", owner(http.HandlerFunc(h.CreateSchedule))).Methods(http.MethodPost)
	r.HandleFunc("/studios/{id}/resources/{rid}/slots", h.GenerateSlots).Methods(http.MethodGet)
	r.Handle("/resources/{id}/exceptions", owner(http.HandlerFunc(h.AddException))).Methods(http.MethodPost)
}