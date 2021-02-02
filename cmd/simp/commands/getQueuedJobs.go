package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getQueuedJobsCallback(cmd *cobra.Command, args []string) {
	fmt.Println("Get all jobs on queue")
}

// GetQueuedJobsCommand is a command to get all jobs on queue
var GetQueuedJobsCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "retorna todos os jobs na fila do STS",
	Short:   "retorna todos os jobs na fila do STS",
	Run:     getQueuedJobsCallback,
	Use:     "getQueuedJobs",
}
