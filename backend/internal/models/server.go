package models

import (
	"context"
	"crypto/tls"
	"github.com/siruspen/logrus"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		TLSConfig:      &tls.Config{ServerName: os.Getenv("DOMAIN")},
	}

	logrus.Printf("Server is listening on address: %s%s", os.Getenv("DOMAIN"), s.httpServer.Addr)

	//return s.httpServer.ListenAndServeTLS("certificate.crt", "private.key")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
