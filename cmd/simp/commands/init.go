package commands

import (
	"os"

	"github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simp/api"
	"github.com/spf13/cobra"
)

var simpCommand = &cobra.Command{}
var client api.ClientInterface

type parameters struct {
	apiClient        api.ClientInterface
	jobID            string
	priority         int
	processName      string
	processParams    []string
	minMemory        int
	minCPU           int
	workingDirectory string
	all              bool
}

var params = parameters{}

// Init faz o parse dos parâmetros e adiciona os comandos
func Init(cl api.ClientInterface) {

	workingDir, _ := os.Getwd()

	CreateJobCommand.Flags().StringVarP(&params.jobID, "id", "i", "no-id", "ID do job que será criado (Inteiro)")
	CreateJobCommand.MarkFlagRequired("id")
	// simpCommand.TraverseChildren = true

	CreateJobCommand.Flags().IntVarP(&params.priority, "priority", "p", 1, "Prioridade do job que será criado (0-1-2)")
	CreateJobCommand.Flags().StringVarP(&params.processName, "name", "n", "", "Caminho absoluto para o binário do job (String)")
	CreateJobCommand.MarkFlagRequired("name")
	CreateJobCommand.MarkFlagFilename("name", "sh")
	CreateJobCommand.Flags().StringArrayVar(&params.processParams, "args", []string{}, "Argumentos do job que será criado (Array de Strings)")
	CreateJobCommand.Flags().IntVarP(&params.minMemory, "mem", "m", 10, "Mínimo de memória disponível para execução do job (0-100)")
	CreateJobCommand.Flags().IntVarP(&params.minCPU, "cpu", "c", 1, "Mínimo de CPU disponível para a execução do job (0-100)")
	CreateJobCommand.Flags().StringVarP(&params.workingDirectory, "work_dir", "w", workingDir, "Diretório de trabalho do job (String)")

	DeleteQueue.Flags().StringVarP(&params.jobID, "id", "i", "no-id", "ID do job que será criado (Inteiro)")
	DeleteQueue.Flags().BoolVarP(&params.all, "all", "a", false, "Aplicar a todos os jobs")
	DeleteExecuting.Flags().StringVarP(&params.jobID, "id", "i", "no-id", "ID do job que será criado (Inteiro)")
	DeleteExecuting.Flags().BoolVarP(&params.all, "all", "a", false, "Aplicar a todos os jobs")
	Delete.AddCommand(&DeleteQueue, &DeleteExecuting)

	GetJobFromQueueCommand.Flags().StringVarP(&params.jobID, "id", "i", "no-id", "ID do job que será criado (Inteiro)")
	GetJobFromQueueCommand.Flags().BoolVarP(&params.all, "all", "a", false, "Aplicar a todos os jobs")
	GetExecutingJobCommand.Flags().StringVarP(&params.jobID, "id", "i", "no-id", "ID do job que será criado (Inteiro)")
	GetExecutingJobCommand.Flags().BoolVarP(&params.all, "all", "a", false, "Aplicar a todos os jobs")
	GetJobs.AddCommand(&GetJobFromQueueCommand, &GetExecutingJobCommand)

	simpCommand.AddCommand(&CreateJobCommand)
	simpCommand.AddCommand(&Delete)
	simpCommand.AddCommand(&GetJobs)

	client = cl
}

// Run executa o comando
func Run() {
	simpCommand.Execute()
}
