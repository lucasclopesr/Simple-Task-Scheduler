package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
)

// HandleGetQueue trata requisitos de m'etodo GET no caminho /queue
func handleGetQueue(w http.ResponseWriter, r *http.Request) {
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

}

// HandleDeleteQueue trata requisitos de m'etodo GET no caminho /queue
func handleDeleteQueue(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	jobID := vars["id_job"]

	err := jh.DeleteJobFromQueue(jobID)
	if err != nil {
		handleError(err, w, r)
		return
	}
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

}

// HandleQueue trata a rota /job/{id_job}
func HandleQueue(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetQueue(w, r)
	case http.MethodDelete:
		handleDeleteQueue(w, r)
	case http.MethodPost:
		handlePostJob(w, r)
	}
}
