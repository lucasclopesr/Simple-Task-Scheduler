package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getJobCallback(cmd *cobra.Command, args []string) {
	fmt.Println("Get job =" + *jobID)
}

// GetJobCommand is a command to get a job
var GetJobCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "get job in the simpd",
	Short:   "get job in the simpd",
	Run:     getJobCallback,
	Use:     "get",
}
