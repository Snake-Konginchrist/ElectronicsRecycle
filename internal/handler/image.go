package handler

import (
	"ElectronicsRecycle/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

func ClassifyImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var request map[string]string
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	base64Image, ok := request["image"]
	if !ok {
		http.Error(w, "Image field is required", http.StatusBadRequest)
		return
	}

	result, err := service.ClassifyImage(base64Image)
	if err != nil {
		log.Printf("Error classifying image: %v", err)
		http.Error(w, "Error classifying image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
