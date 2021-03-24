package memory

import (
	"sync"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/simperr"
)

var lock sync.Mutex
var m memory

type memory map[string]meta.Job
type Memory interface {
	CreateJob(string, meta.Job) error
	DeleteJob(string) error
	GetJob(string) (meta.Job, error)
}

func GetMemory() Memory {
	if m == nil {
		m = make(memory)
	}
	return m
}

// CreateJob adiciona um Job na memória compartilhada de jobs
func (mem memory) CreateJob(id string, job meta.Job) error {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := mem[id]; ok {
		return simperr.NewError().AlreadyExists().Message("job de id " + id + " já existe").Build()
	}
	mem[id] = job
	return nil
}

// DeleteJob remove um Job da memória compartilhada de jobs
func (mem memory) DeleteJob(id string) error {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := mem[id]; !ok {
		return simperr.NewError().Message("job de id " + id + " não encontrado").Build()
	}
	delete(mem, id)
	return nil
}

// GetJob recupera um Job da memória compartilhada de jobs
func (mem memory) GetJob(id string) (meta.Job, error) {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := mem[id]; !ok {
		return meta.Job{}, simperr.NewError().NotFound().Message("job de id " + id + " não encontrado").Build()
	}
	return mem[id], nil
}
