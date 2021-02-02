package handlers

import "github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"

// SetJobHandler assina a variável jh que faz modificações nas estruturas de job
// existentes
func SetJobHandler(jobHandler JobHandler) {
	jh = jobHandler
}

// JobHandler representa uma estrutura que consegue modificar e
// requerir informações sobre os jobs atualmente no sistema
type JobHandler interface {
	CreateJob(string, meta.JobRequest) error
	GetJob(string) (meta.Job, error)
	GetExecutingJobs() ([]meta.Job, error)
	GetQueuedJobs() ([]meta.Job, error)
	DeleteExecutingJobs() error
	DeleteQueuedJobs() error
	DeleteJobFromQueue(string) error
	DeleteExecutingJob(string) error
}

var jh JobHandler
