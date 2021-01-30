package queue

import (
	"container/heap"
	"sync"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/simperr"
)

// lock é um mutex que será usado para garantir a execução correta
// de funções que alteram a estrutura da fila e dos jobs inseridos nela
var lock sync.Mutex

// Len retorna o tamanho da Fila (quantos jobs estão nela)
func (pq *SimpQueueManager) Len() int {
	return len(pq.simpQueue.Queue)
}

// Less compara dois jobs por sua prioridade. Como desejamos que Pop remova o job
// de maior prioridade, utiliza-se o sinal de "maior que"
func (pq *SimpQueueManager) Less(job1Index, job2Index int) bool {
	return pq.simpQueue.Queue[job1Index].Priority > pq.simpQueue.Queue[job2Index].Priority
}

// Swap muda a posição de dois jobs entre si
func (pq *SimpQueueManager) Swap(job1Index, job2Index int) {
	job1ID := pq.simpQueue.Queue[job1Index].ID
	job2ID := pq.simpQueue.Queue[job2Index].ID

	lock.Lock()
	defer lock.Unlock()
	pq.simpQueue.Queue[job1Index], pq.simpQueue.Queue[job2Index] = pq.simpQueue.Queue[job2Index], pq.simpQueue.Queue[job1Index]
	pq.simpQueue.Queue[job1Index].Index = job1Index
	pq.simpQueue.Queue[job2Index].Index = job2Index

	pq.simpQueue.IndexList[job1ID] = job2Index
	pq.simpQueue.IndexList[job2ID] = job1Index
}

// Push insere job na Fila
func (pq *SimpQueueManager) Push(h interface{}) {
	lock.Lock()
	defer lock.Unlock()

	n := len(pq.simpQueue.Queue)
	job := h.(*meta.Job)
	job.Index = n
	pq.simpQueue.Queue = append(pq.simpQueue.Queue, job)
}

// Pop remove o job de maior prioridade da Fila
func (pq *SimpQueueManager) Pop() interface{} {
	lock.Lock()
	defer lock.Unlock()

	old := pq.simpQueue.Queue
	n := len(old)
	job := old[n-1]
	old[n-1] = nil
	job.Index = -1
	pq.simpQueue.Queue = old[0 : n-1]
	return job
}

// GetJobFromQueue recebe o ID de um job e retorna-o, caso esteja na fila. Caso contrario, retorna erro
func (pq *SimpQueueManager) GetJobFromQueue(jobID string) (*meta.Job, error) {
	jobIndex, exists := pq.simpQueue.IndexList[jobID]
	if !exists {
		return nil, simperr.NewError().NotFound().Message("coundn't find job " + jobID + " in queue").Build()
	}
	return pq.simpQueue.Queue[jobIndex], nil
}

// InsertJobIntoQueue insere novo job na fila, caso não exista um com ID igual.
// Caso exista, retorna um erro. Caso contrário, insere o job.
func (pq *SimpQueueManager) InsertJobIntoQueue(job meta.Job) error {
	_, err := pq.GetJobFromQueue(job.ID)
	if err == nil {
		return err
	}
	heap.Push(pq, job)
	return nil
}

// DeleteJobFromQueue remove um job da fila. Caso o job não esteja na fila, retorna um erro
func (pq *SimpQueueManager) DeleteJobFromQueue(jobID string) (interface{}, error) {
	job, err := pq.GetJobFromQueue(jobID)
	if err != nil {
		return nil, err
	}

	lock.Lock()
	defer lock.Unlock()

	removedJob := heap.Remove(pq, job.Index)

	return removedJob, nil
}

// UpdateQueuedJob atualiza as informações de um job que já se encontra na fila. Caso o job não seja
// encontrado na fila OU caso haja uma tentativa de alterar o Index do job, é retornado um erro.
func (pq *SimpQueueManager) UpdateQueuedJob(job meta.Job) error {
	oldJob, err := pq.GetJobFromQueue(job.ID)
	if err != nil {
		return err
	}

	if oldJob.Index != job.Index {
		return simperr.NewError().BadRequest().Message("can't change a job's Index attribute").Build()
	}

	lock.Lock()
	defer lock.Unlock()

	pq.simpQueue.Queue[job.Index].Priority = job.Priority
	pq.simpQueue.Queue[job.Index].ProcessName = job.ProcessName
	pq.simpQueue.Queue[job.Index].ProcessParams = job.ProcessParams

	heap.Fix(pq, job.Index)
	return nil
}
