package meta

// A Job is what compose the priority queue.
type Job struct {
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index         int      // The index of the Job in the heap.
	ID            string   // The identifier of the Job.
	Priority      int      // The priority of the Job in the queue.
	ProcessName   string   // The process' executable name
	ProcessParams []string // Process' inputed/needed parameters
}
