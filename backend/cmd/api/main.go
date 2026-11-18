package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rrspace/backend/internal/cache"
	"github.com/rrspace/backend/internal/config"
	"github.com/rrspace/backend/internal/database"
	"github.com/rrspace/backend/internal/router"
	"github.com/rrspace/backend/internal/storage"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	// PostgreSQL (Neon)
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("database connect error: %v", err)
	}
	defer db.Close()

	// Redis
	redisClient := cache.NewRedisClient(cfg)
	defer redisClient.Close()

	// Cloudflare R2
	r2 := storage.NewR2Client(cfg)

	// Router
	r := router.New(cfg, db, redisClient, r2)

	// Graceful shutdown
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Go API listening on :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Wait for interrupt
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}
	log.Println("server stopped")
}