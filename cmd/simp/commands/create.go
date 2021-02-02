package commands

import (
	"fmt"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
	"github.com/spf13/cobra"
)

func createJobCallback(cmd *cobra.Command, args []string) {
	user := "placeholder"
	err := client.CreateJob(meta.JobRequest{
		User: user,
		Job: meta.Job{
			Priority:         params.priority,
			ID:               params.jobID,
			ProcessName:      params.processName,
			ProcessParams:    params.processParams,
			MinCPU:           params.minCPU,
			MinMemory:        params.minMemory,
			WorkingDirectory: params.workingDirectory,
		},
	}, params.jobID)
	fmt.Printf("Create job = " + params.jobID + "\n")
	fmt.Printf("MinCPU =  %d\n", params.minCPU)
	fmt.Printf("MinMemory =  %d\n", params.minMemory)
	fmt.Println(err)
}

// CreateJobCommand is a command to create a job
var CreateJobCommand = cobra.Command{
	Aliases: []string{"new"},
	Long:    "cria um novo job no STS",
	Short:   "cria um novo job no STS",
	Run:     createJobCallback,
	Use:     "create",
}
