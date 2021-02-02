package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getJobFromQueueCallback(cmd *cobra.Command, args []string) {
	fmt.Println("Get job =" + *jobID)
}

// GetJobFromQueueCommand is a command to get a job
var GetJobFromQueueCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "retorna um job na fila do STS",
	Short:   "retorna um job na fila do STS",
	Run:     getJobFromQueueCallback,
	Use:     "getJobFromQueue",
}
