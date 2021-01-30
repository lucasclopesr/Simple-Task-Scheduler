package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func deleteJobsCallback(cmd *cobra.Command, args []string) {
	fmt.Println("Delete all running jobs")
}

// DeleteJobsCommand is a command to delete all running jobs
var DeleteJobsCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "delete all running jobs in the simpd",
	Short:   "delete all running jobs in the simpd",
	Run:     deleteJobsCallback,
	Use:     "deleteJobs",
}
