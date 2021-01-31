package memory

import (
	"sync"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/simperr"
)

var mem map[string]meta.Job = make(map[string]meta.Job)
var lock sync.Mutex

func CreateJob(id string, job meta.Job) error {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := mem[id]; ok {
		return simperr.NewError().Build()
	}
	mem[id] = job
	return nil
}

func DeleteJob(id string) error {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := mem[id]; !ok {
		return simperr.NewError().Build()
	}
	delete(mem, id)
	return nil
}

func GetJob(id string) (meta.Job, error) {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := mem[id]; !ok {
		return meta.Job{}, simperr.NewError().Build()
	}
	return mem[id], nil
}
