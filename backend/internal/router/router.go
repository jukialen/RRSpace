package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/studiobooking/go-api/internal/cache"
	"github.com/studiobooking/go-api/internal/config"
	"github.com/studiobooking/go-api/internal/handlers"
	"gorm.io/gorm"
)

type R2Client interface {
	UploadFile(ctx any, key string, body any, ct string) (string, error)
}

func New(cfg *config.Config, db *gorm.DB, redis *cache.RedisClient, r2 R2Client) *mux.Router {
	r := mux.NewRouter()

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods(http.MethodGet)

	// Handlers
	authH := handlers.NewAuthHandler()
	studioH := handlers.NewStudioHandler(db)
	resourceH := handlers.NewResourceHandler(db)
	scheduleH := handlers.NewScheduleHandler(db)
	bookingH := handlers.NewBookingHandler(db)

	authH.RegisterRoutes(r)
	studioH.RegisterRoutes(r)
	resourceH.RegisterRoutes(r)
	scheduleH.RegisterRoutes(r)
	bookingH.RegisterRoutes(r)

	return r
}