package gohttpserver

import (
	"fmt"
	"log"
	"net/http"
)

type HttpServer interface {
	ListenAndServe()
	ListenAndServeTLS(certFile, keyFile string)
}

type httpServer struct {
	server http.Server
}

func (s *httpServer) ListenAndServe() {
	fmt.Printf("HTTP server running on address: %v\n", s.server.Addr)
	log.Fatal(s.server.ListenAndServe())
}

func (s *httpServer) ListenAndServeTLS(certFile, keyFile string) {
	fmt.Printf("HTTP server running on address: %v\n", s.server.Addr)
	log.Fatal(s.server.ListenAndServeTLS(certFile, keyFile))
}
