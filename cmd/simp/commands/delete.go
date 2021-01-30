package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func deleteJobCallback(cmd *cobra.Command, args []string) {
	fmt.Println(*jobID)
}

// DeleteJobCommand is a command to delete a job
var DeleteJobCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "delete a job in the simpd",
	Short:   "delete a job in the simpd",
	Run:     deleteJobCallback,
	Use:     "delete",
}
