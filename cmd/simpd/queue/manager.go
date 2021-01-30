package queue

import (
	"container/heap"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
)

// SimpQueueManager define a estrutura que consegue utilizar os métodos
// para manipulação da Fila de Prioridades definidos pela interfaze PQInterface
type SimpQueueManager struct {
	simpQueue *meta.PriorityQueue
}

var simpPQ *SimpQueueManager

// GetQueueManager é a função que retorna a estrutura que manipula a
// Fila de Prioridades. Além disso, garante que exista apenas uma dessa
// estrutura, da mesma forma que a Fila de Prioridades também é única
func GetQueueManager() PQInterface {
	if simpPQ == nil {
		simpPQ = newQueueManager()
		heap.Init(simpPQ)
	}
	return simpPQ
}

func newQueueManager() *SimpQueueManager {
	return &SimpQueueManager{
		simpQueue: &meta.PriorityQueue{
			Queue:     []*meta.Job{},
			IndexList: make(map[string]int),
		},
	}
}
