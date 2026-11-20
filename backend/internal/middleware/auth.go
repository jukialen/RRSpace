package middleware

import (
	"context"
	"net/http"

	"rrspace/backend/internal/models"
)

type contextKey string

const userContextKey contextKey = "user"

// GetUserFromContext zwraca użytkownika z kontekstu
func GetUserFromContext(r *http.Request) *models.User {
	if u, ok := r.Context().Value(userContextKey).(*models.User); ok {
		return u
	}
	return nil
}

func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

// AuthMiddleware — waliduje token Appwrite / JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
			return
		}

		// TODO: walidacja tokena przez Appwrite SDK
		// user, err := appwriteClient.VerifySession(token)
		// if err != nil { ... }

		// Tymczasowo: mock
		user := &models.User{
			ID:    "mock-user-id",
			Role:  "client",
			Email: "mock@example.com",
		}

		ctx := WithUser(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}