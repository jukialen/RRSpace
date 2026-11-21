#!/usr/bin/env bash
set -e

# Funkcja sprzątająca przy zamknięciu skryptu (Ctrl+C)
cleanup() {
    echo ""
    echo "Zatrzymywanie kontenerów Podmana..."
    podman compose -f "podman compose.yml" down
    exit 0
}

# Przechwycenie sygnału przerwania (Ctrl+C) i zakończenia
trap cleanup SIGINT SIGTERM

if [ ! -f "backend/go.sum" ]; then
    echo "Tworzenie backend/go.sum..."
    podman run --rm -v ./backend:/app -w /app golang:1.23-bookworm go mod tidy
fi

echo "Uruchamianie kontenerów (Backend + Redis)..."
podman compose -f "podman compose.yml" up --build -d

echo "Uruchamianie frontendu (Bun)..."
cd frontend && bun run dev