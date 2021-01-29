// Package api é o pacote que define o lado do daemon da API do SIMP. Cria as rotas e inicializa o servidor
// além de fazer os handlers para alterações nos processos. Esses handlers dependem de interfaces
// a serem implementadas, de criação, monitoramento e deleção de jobs.
//
// O valor necessário no corpo dos requests são as estruturas definidas em pkg/meta.
package api

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

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

func (s *Server) serve(ctx context.Context) (err error) {
	server := &http.Server{
		Handler: s.router,
	}
	go func() {
		if err = server.Serve(s.listener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%v", err)
		}
	}()

	log.Println("Listening...")
	<-ctx.Done()
	log.Printf("Gracefully shutting down...")

	ctxShutdown, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer func() {
		cancel()
	}()

	if err = server.Shutdown(ctxShutdown); err != nil {
		log.Fatal("error shutting down server")
	}

	err = os.RemoveAll(transport.UnixSocketAddress)
	if err != nil {
		return err
	}

	log.Println("server shutdown complete")
	if err == http.ErrServerClosed {
		err = nil
	}
	return
}

// Run roda o servidor
func (s *Server) Run() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()
	return s.serve(ctx)
}
