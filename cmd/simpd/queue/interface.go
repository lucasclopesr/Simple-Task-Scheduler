package queue

import (
	"container/heap"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
)

// PQInterface é a interface que define os métodos a serem
// utilizados para manipulação da estrutura Fila de Prioridades
type PQInterface interface {
	heap.Interface
	GetJobFromQueue(jobID string) (meta.Job, error)
	InsertJobIntoQueue(job meta.Job) error
	DeleteJobFromQueue(jobID string) (meta.Job, error)
	UpdateQueuedJob(job meta.Job) error
}
