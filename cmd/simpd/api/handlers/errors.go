package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/simperr"
)

func handleError(err error, w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	var status int
	serr, ok := err.(*simperr.SimpError)
	if !ok {
		status = http.StatusInternalServerError
	} else {
		switch serr.Code {
		case simperr.ErrorAlreadyExists:
			status = http.StatusConflict
		case simperr.ErrorJobLimit, simperr.ErrorMemoryLimit:
			status = http.StatusTooManyRequests
		case simperr.ErrorNotFound:
			status = http.StatusNotFound
		default:
			status = http.StatusInternalServerError
		}
	}

	w.WriteHeader(status)
	encoder.Encode(serr)
}
