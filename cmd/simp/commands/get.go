package commands

import (
	"fmt"

	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/meta"
	"github.com/spf13/cobra"
)

func getJobFromQueueCallback(cmd *cobra.Command, args []string) {
	var err error
	var job meta.Job
	var jobs []meta.Job

	if params.all {
		jobs, err = client.GetQueuedJobs()
		for _, job := range jobs {
			job.PrintJob()
		}
	} else {
		job, err = client.GetJobFromQueue(params.jobID)
		if err == nil {
			job.PrintJob()
		} else {
			fmt.Println(err)
		}
	}
}

func getExecutingJobCallback(cmd *cobra.Command, args []string) {
	var err error
	var job meta.Job
	var jobs []meta.Job

	if params.all {
		jobs, err = client.GetExecutingJobs()
		for _, currJob := range jobs {
			currJob.PrintJob()
		}
	} else {
		job, err = client.GetExecutingJob(params.jobID)
		if err == nil {
			job.PrintJob()
		} else {
			fmt.Println(err)
		}
	}
}

// GetJobs is a command to get jobs from queue or in execution
var GetJobs = cobra.Command{
	Long:  "retorna os detalhes de jobs na fila ou em execução",
	Short: "retorna os detalhes de jobs na fila ou em execução",
	Use:   "get",
}

// GetJobFromQueueCommand is a command to get a job
var GetJobFromQueueCommand = cobra.Command{
	Aliases: []string{"q"},
	Long:    "retorna jobs na fila do STS",
	Short:   "retorna jobs na fila do STS",
	Run:     getJobFromQueueCallback,
	Use:     "queue",
}

// GetExecutingJobCommand is a command to get a job
var GetExecutingJobCommand = cobra.Command{
	Aliases: []string{"e"},
	Long:    "retorna jobs em execução no STS",
	Short:   "retorna jobs em execução no STS",
	Run:     getExecutingJobCallback,
	Use:     "exec",
}
