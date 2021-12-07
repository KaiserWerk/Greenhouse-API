package handler

import (
	"encoding/json"
	"net/http"

	"github.com/KaiserWerk/Greenhouse-Manager/internal/entity"
	"github.com/KaiserWerk/Greenhouse-Manager/internal/influx"
)

func (h HttpHandler) ReceiveHandler(w http.ResponseWriter, r *http.Request) {
	var m entity.Measurement
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "could not decode JSON", http.StatusBadRequest)
	}
	_ = r.Body.Close()

	influx.InsertMeasurement(m)
	w.WriteHeader(http.StatusCreated)
}
