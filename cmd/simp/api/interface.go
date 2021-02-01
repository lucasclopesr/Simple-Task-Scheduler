package api

import "github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"

// ClientInterface is the interface for interacting with the simp daemon restful API
type ClientInterface interface {
	CreateJob(request meta.JobRequest, id string) error
	DeleteJobFromQueue(id string) error
	DeleteExecutingJob(id string) error
	GetJobFromQueue(id string) (meta.Job, error)
	GetExecutingJob(id string) (meta.Job, error)
	GetExecutingJobs() ([]meta.Job, error)
	GetQueuedJobs() ([]meta.Job, error)
	DeleteQueue() error
	DeleteExecutingJobs() error
}
