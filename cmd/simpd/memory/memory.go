package memory

import (
	"sync"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/simperr"
)

var mem map[string]meta.Job = make(map[string]meta.Job)
var lock sync.Mutex

// CreateJob adiciona um Job na memória compartilhada de jobs
func CreateJob(id string, job meta.Job) error {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := mem[id]; ok {
		return simperr.NewError().AlreadyExists().Message("job de id " + id + " já existe").Build()
	}
	mem[id] = job
	return nil
}

// DeleteJob remove um Job da memória compartilhada de jobs
func DeleteJob(id string) error {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := mem[id]; !ok {
		return simperr.NewError().Message("job de id " + id + " não encontrado").Build()
	}
	delete(mem, id)
	return nil
}

// GetJob recupera um Job da memória compartilhada de jobs
func GetJob(id string) (meta.Job, error) {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := mem[id]; !ok {
		return meta.Job{}, simperr.NewError().NotFound().Message("job de id " + id + " não encontrado").Build()
	}
	return mem[id], nil
}
