package handlers

import (
	"log"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/simperr"
)

// NewJobHandlerMock creates a mock of a job handler for server testing
func NewJobHandlerMock() JobHandler {
	return &jobHandlerMock{}
}

type jobHandlerMock map[string]meta.Job

// CreateJob mock
func (j jobHandlerMock) CreateJob(s string, jr meta.JobRequest) error {
	if _, ok := j[s]; ok {
		return &simperr.SimpError{
			Code:    403,
			Message: "already exists",
		}
	}
	j[s] = jr.Job
	return nil
}

// DeleteJob mock
func (j jobHandlerMock) DeleteJob(s string) error {
	if _, ok := j[s]; !ok {
		return &simperr.SimpError{
			Code:    403,
			Message: "job not found",
		}
	}
	delete(j, s)
	return nil
}

// GetJob mock
func (j jobHandlerMock) GetJob(s string) (meta.Job, error) {
	job, ok := j[s]
	if !ok {
		return meta.Job{}, &simperr.SimpError{
			Code:    403,
			Message: "job not found",
		}
	}
	return job, nil
}

// GetExecutingJobs mock
func (j jobHandlerMock) GetExecutingJobs() ([]meta.Job, error) {
	ret := []meta.Job{}
	for _, v := range j {
		ret = append(ret, v)
	}
	return ret, nil
}

// DeleteExecutingJobs mock
func (j *jobHandlerMock) DeleteExecutingJobs() error {
	log.Println("delete executing")
	return nil
}

// GetQueuedJobs mock
func (j jobHandlerMock) GetQueuedJobs() ([]meta.Job, error) {
	ret := []meta.Job{}
	for _, v := range j {
		ret = append(ret, v)
	}
	return ret, nil
}

// DeleteQueuedJobs mock
func (j *jobHandlerMock) DeleteQueuedJobs() error {
	log.Println("delete queued")
	return nil
}
