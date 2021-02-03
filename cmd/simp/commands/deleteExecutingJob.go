package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func deleteExecutingJobCallback(cmd *cobra.Command, args []string) {
	fmt.Println(*jobID)
}

// DeleteExecutingJobCommand is a command to delete a job
var DeleteExecutingJobCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "deleta um job em execução no STS",
	Short:   "deleta um job em execução no STS",
	Run:     deleteExecutingJobCallback,
	Use:     "deleteExecutingJob",
}
