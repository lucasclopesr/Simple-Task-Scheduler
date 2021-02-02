package jobhandler

import (
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/api/handlers"
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/memory"
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/processes"
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/queue"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/simperr"
)

// NewJobHandler cria um job handler para alocar recursos para os jobs
func NewJobHandler() handlers.JobHandler {
	return &jobHandler{}
}

type jobHandler struct{}

// CreateJob valida um job e o insere na fila de prioridades
func (j jobHandler) CreateJob(s string, jr meta.JobRequest) error {
	if _, err := memory.GetJob(s); err == nil {
		return simperr.NewError().AlreadyExists().Build()
	}
	queue := queue.GetQueueManager()
	err := queue.InsertJobIntoQueue(jr.Job)
	if queue.Len() == 1 {
		processes.GetProcessManager().CreateFirstJob()
	}
	memory.CreateJob(s, jr.Job)
	return err
}

// DeleteExecutingJob deleta job da fila de prioridades, caso exista. Caso contrário,
// retorna um erro.
func (j jobHandler) DeleteJobFromQueue(s string) error {
	_, err := queue.GetQueueManager().DeleteJobFromQueue(s)
	if err != nil {
		return err
	}
	memory.DeleteJob(s)
	return nil
}

// GetJob retorna job com o ID dado. Caso o job não exista, retorna erro
func (j jobHandler) GetJob(jobID string) (job meta.Job, err error) {

	if job, err = memory.GetJob(jobID); err != nil {
		return job, simperr.NewError().DoesNotExist().Build()
	}
	return job, err
}

// GetExecutingJobs retorna todos os jobs em execução
func (j jobHandler) GetExecutingJobs() ([]meta.Job, error) {
	ret := []meta.Job{}
	return ret, nil // Todo: implement
}

// DeleteExecutingJobs deleta todos os jobs em execução
func (j *jobHandler) DeleteExecutingJobs() error {
	return nil // Todo: implement
}

// GetQueuedJobs retorna todos os jobs que se encontram na Fila de Prioridades.
// Caso a fila esteja vazia, retorna um erro.
func (j jobHandler) GetQueuedJobs() ([]meta.Job, error) {
	ret, err := queue.GetQueueManager().ReturnAllQueuedJobs()

	if err != nil {
		return nil, err
	}
	return ret, nil
}

// DeleteQueuedJobs deleta todos os jobs da fila de prioridades
func (j *jobHandler) DeleteQueuedJobs() error {
	return nil // Todo: implement
}

// DeleteExecutingJob deleta job em execução, caso exista. Caso contrário,
// retorna um erro.
func (j jobHandler) DeleteExecutingJob(jobID string) error {
	err := processes.GetProcessManager().DeleteJob(jobID)
	if err != nil {
		return err
	}
	memory.DeleteJob(jobID)
	return nil
}
