package commands

import (
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simp/api"
	"github.com/spf13/cobra"
)

var apiClient api.ClientInterface
var user *string
var id *string
var minMemory *int
var minCPU *int
var jobArgs *[]string

var simpCommand = &cobra.Command{}
var client api.ClientInterface

func Init(cl api.ClientInterface) {
	user = CreateJobCommand.Flags().StringP("user", "u", "", "user to create the job by")
	id = CreateJobCommand.Flags().StringP("job_id", "i", "no-id", "the id for the job that will be created")
	minMemory = CreateJobCommand.Flags().IntP("mem", "m", 300, "the id for the job that will be created")
	minCPU = CreateJobCommand.Flags().IntP("cpu", "c", 200, "the id for the job that will be created")
	jobArgs = CreateJobCommand.Flags().StringArrayP("args", "a", []string{}, "arguments for the job to be created")

	simpCommand.AddCommand(&CreateJobCommand)
	client = cl
}

func Run() {
	simpCommand.Execute()
}
