package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getJobsCallback(cmd *cobra.Command, args []string) {
	fmt.Println("Get all jobs")
}

// GetJobsCommand is a command to get all running jobs
var GetJobsCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "get all running jobs in the simpd",
	Short:   "get all running jobs in the simpd",
	Run:     getJobsCallback,
	Use:     "getJobs",
}
