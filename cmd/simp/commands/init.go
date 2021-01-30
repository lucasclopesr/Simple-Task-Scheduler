package commands

import (
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simp/api"
	"github.com/spf13/cobra"
)

var apiClient api.ClientInterface

var jobID *string

var priority *int

var processName *string

var processParams *[]string

var minMemory *int

var minCPU *int

var simpCommand = &cobra.Command{}

var client api.ClientInterface

// Init faz o parse dos parâmetros
func Init(cl api.ClientInterface) {

	jobID = CreateJobCommand.Flags().StringP("job_id", "i", "no-id", "the id for the job that will be created")
	priority = CreateJobCommand.Flags().IntP("priority", "p", 1, "the priority for the job that will be created")
	processName = CreateJobCommand.Flags().StringP("name", "n", "", "the path for the job that will be created")
	processParams = CreateJobCommand.Flags().StringArrayP("args", "a", []string{}, "arguments for the job to be created")
	minMemory = CreateJobCommand.Flags().IntP("mem", "m", 300, "the id for the job that will be created")
	minCPU = CreateJobCommand.Flags().IntP("cpu", "c", 200, "the id for the job that will be created")

	simpCommand.AddCommand(&CreateJobCommand)
	simpCommand.AddCommand(&DeleteJobCommand)
	simpCommand.AddCommand(&GetJobCommand)

	client = cl
}

// Run executa o comando
func Run() {
	simpCommand.Execute()
}
