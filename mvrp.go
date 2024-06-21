package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Server struct {
	addr     string
	certFile string
	keyFile  string
}

func NewServer(addr, certFile, keyFile string) *Server {
	return &Server{addr: addr, certFile: certFile, keyFile: keyFile}
}

func (s *Server) Run() {
	// Load server certificate and key
	cert, err := tls.LoadX509KeyPair(s.certFile, s.keyFile)
	if err != nil {
		log.Fatalf("Failed to load key pair: %v", err)
	}

	// Load client CA certificate
	caCert, err := ioutil.ReadFile("path/to/ca-cert.pem")
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Configure TLS settings
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	// Create HTTPS server
	server := &http.Server{
		Addr:      s.addr,
		TLSConfig: tlsConfig,
		Handler:   http.HandlerFunc(s.handleRequest),
	}

	log.Printf("Starting server on %s", s.addr)
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)
	w.Header().Set("X-Custom-Header", "EncryptedValue")

	switch r.Method {
	case "OPTIONS":
		w.Header().Set("Allow", "OPTIONS, CREATE, READ, EMIT, BURN")
		w.WriteHeader(http.StatusNoContent)
	case "CREATE":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Resource created\n"))
	case "READ":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Resource read\n"))
	case "EMIT":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Event emitted\n"))
	case "BURN":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Resource burned\n"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed\n"))
	}
}
