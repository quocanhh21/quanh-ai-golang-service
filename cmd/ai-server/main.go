package main

import (
    "log"
    "net/http"
    "github.com/joho/godotenv"

    "github.com/quocanhh21/quanh-ai-golang-service/internal/handler"
)

func main() {
    // Load .env file if present
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, relying on system environment")
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/api/ai/analyze", handler.AnalyzeEssayHandler)

    log.Println("🚀 Golang AI Service running on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal("Server error:", err)
    }
}