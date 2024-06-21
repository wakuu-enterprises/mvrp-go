package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	addr     string
	tlsConfig *tls.Config
}

func NewClient(addr, keyFile, certFile, caFile string) *Client {
	// Load client certificate and key
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load key pair: %v", err)
	}

	// Load CA certificate
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Configure TLS settings
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	return &Client{addr: addr, tlsConfig: tlsConfig}
}

func (c *Client) SendRequest(method, url, body string) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: c.tlsConfig,
		},
	}

	req, err := http.NewRequest(method, "https://"+c.addr+url, strings.NewReader(body))
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}
