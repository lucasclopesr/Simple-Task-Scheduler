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
	Long:    "retorna um job em execução no STS",
	Short:   "retorna um job em execução no STS",
	Run:     getExecutingJobCallback,
	Use:     "getExecutingJob",
}
