package handler

import (
	"encoding/json"
	"github.com/KaiserWerk/Greenhouse-Manager/internal/storage"
	"net/http"

	"github.com/KaiserWerk/Greenhouse-Manager/internal/entity"
)

func (h HttpHandler) ReceiveHandler(w http.ResponseWriter, r *http.Request) {
	var m entity.Measurement
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "could not decode JSON", http.StatusBadRequest)
	}
	_ = r.Body.Close()

	storage.WriteMeasurement(&m)
	w.WriteHeader(http.StatusCreated)
}
