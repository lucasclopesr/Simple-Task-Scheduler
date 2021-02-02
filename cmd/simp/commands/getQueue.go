package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getQueueCallback(cmd *cobra.Command, args []string) {
	fmt.Println("Get all jobs on queue")
}

// GetQueueCommand is a command to get all jobs on queue
var GetQueueCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "get all jobs on queue in the simpd",
	Short:   "get all jobs on queue in the simpd",
	Run:     getQueueCallback,
	Use:     "getQueue",
}
