package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getExecutingJobsCallback(cmd *cobra.Command, args []string) {
	fmt.Println("Get all running jobs")
}

// GetExecutingJobsCommand is a command to get all running jobs
var GetExecutingJobsCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "retorna todos os jobs em execução no STS",
	Short:   "retorna todos os jobs em execução no STS",
	Run:     getExecutingJobsCallback,
	Use:     "getExecutingJobs",
}
