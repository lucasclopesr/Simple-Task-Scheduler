package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func deleteJobFromQueueCallback(cmd *cobra.Command, args []string) {
	fmt.Println(*jobID)
}

// DeleteJobFromQueueCommand is a command to delete a job
var DeleteJobFromQueueCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "deleta job na fila do STS",
	Short:   "deleta job na fila do STS",
	Run:     deleteJobFromQueueCallback,
	Use:     "deleteJobFromQueue",
}
