package meta

// Job é a estrutura que contém as informações do job a ser executado pelo SO e
// o index e prioridade referentes à sua posição na Fila de Prioridades
type Job struct {
	Index             int // Index do Job no heap (Fila de Prioridades)
	ID                string
	Priority          int      // Prioridade de execução do Job
	ProcessName       string   // Nome do processo a ser executado
	ProcessParams     []string // Parâmetros para execução de um processo
	MinCPU, MinMemory int
	WorkingDirectory  string
}
