package handler

import (
	"encoding/json"
	"net/http"

	"github.com/quocanhh21/quanh-ai-golang-service/internal/model"
	"github.com/quocanhh21/quanh-ai-golang-service/internal/service"
)

func AnalyzeEssayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var input model.AnalyzeRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := service.AnalyzeWithOpenAI(input.Text)
	if err != nil {
		http.Error(w, "AI error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
