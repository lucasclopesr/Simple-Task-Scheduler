package queue

import (
	"sort"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
)

// PQInterface é a interface que define os métodos a serem
// utilizados para manipulação da estrutura Fila de Prioridades
type PQInterface interface {
	sort.Interface
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Push(x interface{})
	Pop() interface{}
	// Remove(h Interface, i int) interface{}
	GetJobFromQueue(jobID string) (*meta.Job, error)
	InsertJobIntoQueue(job meta.Job) error
	DeleteJobFromQueue(jobID string) (interface{}, error)
	UpdateQueuedJob(job meta.Job) error
}
