package jobhandler

import (
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/api/handlers"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
)

// NewJobHandler creates a job handler to allocate resources for jobs
func NewJobHandler() handlers.JobHandler {
	return &jobHandler{}
}

type jobHandler map[string]meta.Job

// CreateJob mock
func (j jobHandler) CreateJob(s string, jr meta.JobRequest) error {
	/* Needs to:
	*	Validate job and arguments
	* Insert into queue (or send directly to execution)
	 */
	return nil // Todo: implement
}

// DeleteJob mock
func (j jobHandler) DeleteJob(s string) error {
	return nil // Todo: implement
}

// GetJob mock
func (j jobHandler) GetJob(s string) (meta.Job, error) {
	job := j[s]
	return job, nil // Todo: implement
}

// GetExecutingJobs mock
func (j jobHandler) GetExecutingJobs() ([]meta.Job, error) {
	ret := []meta.Job{}
	return ret, nil // Todo: implement
}

// DeleteExecutingJobs mock
func (j *jobHandler) DeleteExecutingJobs() error {
	return nil // Todo: implement
}

// GetQueuedJobs mock
func (j jobHandler) GetQueuedJobs() ([]meta.Job, error) {
	ret := []meta.Job{}
	return ret, nil // Todo: implement
}

// DeleteQueuedJobs mock
func (j *jobHandler) DeleteQueuedJobs() error {
	return nil // Todo: implement
}
