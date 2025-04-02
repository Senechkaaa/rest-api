package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        handler,
		//	Handler определяет роутинг
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
	/*
		s.httpServer.Shutdown(ctx): Метод Shutdown завершает работу сервера с заданным контекстом. Это аккуратно завершает все активные соединения и завершает работу сервера.
	*/
}
