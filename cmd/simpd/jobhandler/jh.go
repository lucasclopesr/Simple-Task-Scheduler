package jobhandler

import (
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/api/handlers"
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/memory"
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/processes"
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/queue"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/simperr"
)

// NewJobHandler creates a job handler to allocate resources for jobs
func NewJobHandler() handlers.JobHandler {
	return &jobHandler{}
}

type jobHandler struct{}

// CreateJob validates job and inserts into queue
func (j jobHandler) CreateJob(s string, jr meta.JobRequest) error {
	if _, err := memory.GetJob(s); err == nil {
		return simperr.NewError().Build()
	}
	queue := queue.GetQueueManager()
	err := queue.InsertJobIntoQueue(jr.Job)
	if queue.Len() == 1 {
		processes.GetProcessManager().CreateFirstJob()
	}
	memory.CreateJob(s, jr.Job)
	return err
}

// DeleteJob deletes a job from queue
func (j jobHandler) DeleteJob(s string) error {
	if _, err := memory.GetJob(s); err != nil {
		return simperr.NewError().Build()
	}
	queue := queue.GetQueueManager()
	_, err := queue.DeleteJobFromQueue(s)
	memory.DeleteJob(s)
	return err
}

// GetJob finds job of ID s and returns in the format meta.Job
func (j jobHandler) GetJob(s string) (job meta.Job, err error) {

	if job, err = memory.GetJob(s); err != nil {
		return job, simperr.NewError().Build()
	}
	return job, err
}

// GetExecutingJobs returns all currently executing jobs in the format meta.job
func (j jobHandler) GetExecutingJobs() ([]meta.Job, error) {
	ret := []meta.Job{}
	return ret, nil // Todo: implement
}

// DeleteExecutingJobs deletes all currently executing jobs
func (j *jobHandler) DeleteExecutingJobs() error {
	return nil // Todo: implement
}

// GetQueuedJobs returns all jobs currently in queue
func (j jobHandler) GetQueuedJobs() ([]meta.Job, error) {
	ret := []meta.Job{}
	return ret, nil // Todo: implement
}

// DeleteQueuedJobs deletes all jobs currently in queue
func (j *jobHandler) DeleteQueuedJobs() error {
	return nil // Todo: implement
}
