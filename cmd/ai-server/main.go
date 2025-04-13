package main

import (
    "log"
    "net/http"

    "github.com/yourname/quanh-ai-golang-service/internal/handler"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/ai/analyze", handler.AnalyzeEssayHandler)

    log.Println("🚀 Golang AI Service running on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal("Server error:", err)
    }
}