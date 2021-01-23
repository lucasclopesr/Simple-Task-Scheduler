package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
)

// HandleGetJob trata requisitos de m'etodo GET no caminho /job
func handleGetJob(w http.ResponseWriter, r *http.Request) {
	var job meta.Job
	encoder := json.NewEncoder(w)

	vars := mux.Vars(r)
	jobID := vars["id_job"]

	job, err := jh.GetJob(jobID)
	if err != nil {
		handleError(err, w, r)
		return
	}
	err = encoder.Encode(job)
	if err != nil {
		handleError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// HandlePostJob trata requisitos de m'etodo POST no caminho /job
func handlePostJob(w http.ResponseWriter, r *http.Request) {
	var job meta.JobRequest

	vars := mux.Vars(r)
	jobID := vars["id_job"]

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&job)
	if err != nil {
		handleError(err, w, r)
		return
	}

	err = jh.CreateJob(jobID, job)
	if err != nil {
		handleError(err, w, r)
	}

	w.WriteHeader(http.StatusOK)
}

// HandleDeleteJob trata requisitos de m'etodo DELETE no caminho /job
func handleDeleteJob(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	jobID := vars["id_job"]

	err := jh.DeleteJob(jobID)
	if err != nil {
		handleError(err, w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// HandleJob trata a rota /job/{id_job}
func HandleJob(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetJob(w, r)
	case http.MethodPost:
		handlePostJob(w, r)
	case http.MethodDelete:
		handleDeleteJob(w, r)
	}
}
