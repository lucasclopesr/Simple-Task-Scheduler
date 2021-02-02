package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func deleteExecutingJobsCallback(cmd *cobra.Command, args []string) {
	fmt.Println("Delete all running jobs")
}

// DeleteExecutingJobsCommand is a command to delete all running jobs
var DeleteExecutingJobsCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "deleta todos os jobs em execução no STS",
	Short:   "deleta todos os jobs em execução no STS",
	Run:     deleteExecutingJobsCallback,
	Use:     "deleteExecutingJobs",
}
