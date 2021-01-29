package api

import "github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"

// ClientInterface is the interface for interacting with the simp daemon restful API
type ClientInterface interface {
	CreateJob(request meta.JobRequest, id string) error
	DeleteJob(id string) error
	GetJob(id string) (meta.Job, error)
	GetExecutingJobs() ([]meta.Job, error)
	GetQueuedJobs() ([]meta.Job, error)
	DeleteQueue() error
	DeleteExecutingJobs() error
}
