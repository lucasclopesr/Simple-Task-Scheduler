package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
)

// HandleGetQueue trata requisitos de m'etodo GET no caminho /queue
func handleGetQueued(w http.ResponseWriter, r *http.Request) {
	var jobs []meta.Job
	encoder := json.NewEncoder(w)

	jobs, err := jh.GetQueuedJobs()
	if err != nil {
		handleError(err, w, r)
		return
	}
	err = encoder.Encode(jobs)
	if err != nil {
		handleError(err, w, r)
		return
	}

}

// HandleDeleteQueue trata requisitos de m'etodo GET no caminho /queue
func handleDeleteQueued(w http.ResponseWriter, r *http.Request) {

	err := jh.DeleteQueuedJobs()
	if err != nil {
		handleError(err, w, r)
		return
	}
}

// HandleQueued trata a rota /job/{id_job}
func HandleQueued(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetQueue(w, r)
	case http.MethodDelete:
		handleDeleteQueue(w, r)
	}
}
