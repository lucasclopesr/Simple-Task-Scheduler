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
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Job %s criado\n", params.jobID)
	}
}

// CreateJobCommand is a command to create a job
var CreateJobCommand = cobra.Command{
	Aliases: []string{"new", "run"},
	Long:    "cria um novo job no STS",
	Short:   "cria um novo job no STS",
	Run:     createJobCallback,
	Use:     "create",
}
