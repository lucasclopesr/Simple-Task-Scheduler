package meta

// PriorityQueue implementa o tipo heap.Interface que guarda Jobs
type PriorityQueue struct {
	Queue     []*Job
	IndexList map[string]int
}
