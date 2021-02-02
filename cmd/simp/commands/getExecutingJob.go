package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getExecutingJobCallback(cmd *cobra.Command, args []string) {
	fmt.Println("Get job =" + *jobID)
}

// GetExecutingJobCommand is a command to get a job
var GetExecutingJobCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "get executing job in the simpd",
	Short:   "get executing job in the simpd",
	Run:     getExecutingJobCallback,
	Use:     "getExecutingJob",
}
