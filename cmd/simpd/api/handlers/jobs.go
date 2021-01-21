package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
)

// HandleGetJobs trata requisitos de m'etodo GET no caminho /jobs
func handleGetJobs(w http.ResponseWriter, r *http.Request) {
	var jobs []meta.Job
	encoder := json.NewEncoder(w)

	jobs, err := jh.GetExecutingJobs()
	if err != nil {
		handleError(err, w, r)
		return
	}
	err = encoder.Encode(jobs)
	if err != nil {
		handleError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// HandleDeleteQueue trata requisitos de m'etodo GET no caminho /queue
func handleDeleteJobs(w http.ResponseWriter, r *http.Request) {

	err := jh.DeleteExecutingJobs()
	if err != nil {
		handleError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// HandleJobs trata a rota /jobs
func HandleJobs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetJobs(w, r)
	case http.MethodDelete:
		handleDeleteJobs(w, r)
	}
}
