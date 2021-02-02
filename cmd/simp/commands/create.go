package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func createJobCallback(cmd *cobra.Command, args []string) {
	fmt.Printf("Create job = " + *jobID + "\n")
	fmt.Printf("MinCPU =  %d\n", *minCPU)
	fmt.Printf("MinMemory =  %d\n", *minMemory)
}

// CreateJobCommand is a command to create a job
var CreateJobCommand = *&cobra.Command{
	Aliases: []string{"new"},
	Long:    "cria um novo job no STS",
	Short:   "cria um novo job no STS",
	Run:     createJobCallback,
	Use:     "create",
}
