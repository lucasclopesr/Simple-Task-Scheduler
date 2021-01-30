package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func createJobCallback(cmd *cobra.Command, args []string) {
	fmt.Println(*id)
	fmt.Println(*minCPU)
	fmt.Println(*minMemory)
	fmt.Println(*user)
	fmt.Println(*jobArgs)
}

// CreateJobCommand is a command to create a job
var CreateJobCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "create a new job in the simpd",
	Short:   "create a new job in the simpd",
	Run:     createJobCallback,
	Use:     "create",
}
