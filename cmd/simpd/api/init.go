package api

import "github.com/lucasclopesr/Simple-Task-Scheduler/cmd/simpd/api/handlers"

// Init inicializa o server com as rotas necessárias e o handler que faz alterações nos
// processos existentes
func (s *Server) Init(jobHandler handlers.JobHandler) {
	s.router.HandleFunc("/job/{id_job}",
		handlers.JSON(handlers.HandleJob),
	).Methods("GET", "POST", "DELETE")

	s.router.HandleFunc("/jobs",
		handlers.JSON(handlers.HandleJobs),
	).Methods("GET", "DELETE")

	s.router.HandleFunc("queue",
		handlers.JSON(handlers.HandleQueue),
	).Methods("GET", "DELETE")

	// especifica o objeto que fará alteração nos jobs e processos
	handlers.SetJobHandler(jobHandler)
}
