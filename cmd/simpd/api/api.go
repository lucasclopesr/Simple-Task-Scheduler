// Package api é o pacote que define o lado do daemon da API do SIMP. Cria as rotas e inicializa o servidor
// além de fazer os handlers para alterações nos processos. Esses handlers dependem de interfaces
// a serem implementadas, de criação, monitoramento e deleção de jobs.
//
// O valor necessário no corpo dos requests são as estruturas definidas em pkg/meta.
package api

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasclopesr/Simple-Task-Scheduler/pkg/transport"
)

// Server é um servidor que trata as rotas e requests da API rest entre
// CLI e daemon do SIMP
type Server struct {
	router   *mux.Router
	listener net.Listener
}

// NewServer inicializa um servidor com a configuração de unix socket
func NewServer() (Server, error) {
	// ouvindo em unix sockets
	listener, err := net.Listen("unix", transport.UnixSocketAddress)
	if err != nil {
		return Server{}, err
	}

	return Server{
		listener: listener,
		router:   mux.NewRouter(),
	}, nil
}

// Run roda o servidor
func (s *Server) Run() error {
	return http.Serve(s.listener, s.router)
}
