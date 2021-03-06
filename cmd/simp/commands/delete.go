package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func deleteJobFromQueueCallback(cmd *cobra.Command, args []string) {
	var err error

	if params.all {
		err = client.DeleteQueue()
	} else {
		err = client.DeleteJobFromQueue(params.jobID)
	}

	if err == nil {
		if params.all {
			fmt.Println("Todos os jobs da fila foram removidos")
		} else if params.jobID == "" {
			fmt.Println("Parâmetro all não setado e job id não identificado.")
		} else {
			fmt.Printf("Job de ID " + params.jobID + " removido da fila\n")
		}
	} else {
		fmt.Println(err)
	}
}

func deleteExecutingJobCallback(cmd *cobra.Command, args []string) {
	var err error

	if params.all {
		err = client.DeleteExecutingJobs()
	} else {
		err = client.DeleteExecutingJob(params.jobID)
	}

	if err == nil {
		if params.all {
			fmt.Println("Todos os jobs em execução foram cancelados")
		} else if params.jobID == "" {
			fmt.Println("Parâmetro all não setado e job id não identificado.")
		} else {
			fmt.Printf("Job de ID " + params.jobID + " cancelado\n")
		}
	} else {
		fmt.Println(err)
	}
}

// Delete is a command to delete a job
var Delete = cobra.Command{
	Aliases: []string{"del", "remove"},
	Long:    "deleta jobs na fila ou em execução",
	Short:   "deleta jobs na fila ou em execução",
	Use:     "delete",
}

// DeleteQueue is a command to delete a job
var DeleteQueue = cobra.Command{
	Aliases: []string{"q"},
	Long:    "deleta jobs na fila do STS",
	Short:   "deleta jobs na fila do STS",
	Run:     deleteJobFromQueueCallback,
	Use:     "queue",
}

// DeleteExecuting is a command to delete a job
var DeleteExecuting = cobra.Command{
	Aliases: []string{"e", "executing", "running"},
	Long:    "deleta jobs em execução",
	Short:   "deleta jobs execução",
	Run:     deleteExecutingJobCallback,
	Use:     "exec",
}
