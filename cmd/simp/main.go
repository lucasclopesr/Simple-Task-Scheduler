package main

import (
	"fmt"

	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simp/api"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
)

func main() {
	cl := api.NewClient()
	fmt.Println(cl.CreateJob(meta.JobRequest{
		User: "pdrim",
		Job: meta.Job{
			MinCPU:           100,
			MinMemory:        100,
			Priority:         5,
			ID:               "1234",
			ProcessParams:    []string{"script.sh"},
			ProcessName:      "/bin/bash",
			WorkingDirectory: "/home/doc/workdir/Simple-Task-Scheduler/scripts",
		},
	}, "1234"))
	fmt.Println(cl.GetJob("1234"))
}
