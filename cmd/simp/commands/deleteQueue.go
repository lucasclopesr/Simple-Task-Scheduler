package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func deleteQueueCallback(cmd *cobra.Command, args []string) {
	fmt.Println("Delete all jobs on queue")
}

// DeleteQueueCommand is a command to delete all jobs on queue
var DeleteQueueCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "deleta todos os jobs na fila do STS",
	Short:   "deleta todos os jobs na fila do STS",
	Run:     deleteQueueCallback,
	Use:     "deleteQueue",
}
