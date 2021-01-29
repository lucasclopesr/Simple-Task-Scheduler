package meta

// A PriorityQueue implements heap.Interface and holds Jobs.
type PriorityQueue struct {
	Queue     []*Job
	IndexList map[string]int
}
