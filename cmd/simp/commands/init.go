package commands

import (
	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simp/api"
	"github.com/spf13/cobra"
)

var apiClient api.ClientInterface

var jobID *string

var priority *int

var processName *string

var processParams *[]string

var minMemory *int

var minCPU *int

var simpCommand = &cobra.Command{}

var workingDirectory *string

var client api.ClientInterface

// Init faz o parse dos parâmetros e adiciona os comandos
func Init(cl api.ClientInterface) {

	jobID = simpCommand.Flags().StringP("job_id", "i", "no-id", "ID do job que será criado (Inteiro)")
	priority = simpCommand.Flags().IntP("priority", "p", 1, "Prioridade do job que será criado (0-1-2)")
	processName = simpCommand.Flags().StringP("name", "n", "", "Caminho absoluto para o binário do job (String)")
	processParams = simpCommand.Flags().StringArrayP("args", "a", []string{}, "Argumentos do job que será criado (Array de Strings)")
	minMemory = simpCommand.Flags().IntP("mem", "m", 50, "Mínimo de memória disponível para execução do job (0-100)")
	minCPU = simpCommand.Flags().IntP("cpu", "c", 50, "Mínimo de CPU disponível para a execução do job (0-100)")
	workingDirectory = simpCommand.Flags().StringP("work_dir", "w", "", "Diretório de trabalho do job (String)")

	simpCommand.AddCommand(&CreateJobCommand)

	simpCommand.AddCommand(&DeleteJobFromQueueCommand)
	simpCommand.AddCommand(&DeleteExecutingJobCommand)

	simpCommand.AddCommand(&GetJobFromQueueCommand)
	simpCommand.AddCommand(&GetExecutingJobCommand)

	simpCommand.AddCommand(&GetExecutingJobsCommand)
	simpCommand.AddCommand(&GetQueuedJobsCommand)

	simpCommand.AddCommand(&DeleteExecutingJobsCommand)
	simpCommand.AddCommand(&DeleteQueueCommand)

	client = cl
}

// Run executa o comando
func Run() {
	simpCommand.Execute()
}
