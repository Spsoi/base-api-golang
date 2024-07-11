package api

import (
	"Azgart/internal/interfaces/http"
	"log"
	stdhttp "net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := stdhttp.NewServeMux()

	http.RegisterHandlers(router)

	v1 := stdhttp.NewServeMux()

	v1.Handle("/api/v1/", stdhttp.StripPrefix("/api/v1", router))

	middlewareChain := http.MiddlewareChain(
		http.RequestLoggerMiddleware,
		http.RequireAuthMiddleware,
	)

	server := &stdhttp.Server{
		Addr:    s.addr,
		Handler: middlewareChain(v1),
	}

	log.Printf("Server has started on %s", s.addr)

	return server.ListenAndServe()
}
