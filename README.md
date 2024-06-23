# Muvor Protocol (MVRP)

## Description

A custom protocol implementation for Muvor Protocol (MVRP) with modern features.

## Installation

```bash
go get -v github.com/wakuu-enterprises/mvrp-go
```

```bash
go mod tidy
```

## Client

```bash
package main

import (
	"fmt"
	"log"
	"mvrp/client"
)

func main() {
	cl := client.NewClient("127.0.0.1:8443", "path/to/client-key.pem", "path/to/client-cert.pem", "path/to/ca-cert.pem")
	response, err := cl.SendRequest("CREATE", "/", "Hello, secure server!")
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	fmt.Println("Response:", response)
}
```

## Video

```bash
package main

import (
	"mvrp/mvvp"
)

func main() {
	video.ProcessSegments("/path/to/uploads", "/path/to/structured")
}
```

## Server

```bash
package main

import (
	"mvrp/mvrp"
)

func main() {
	srv := server.NewServer("127.0.0.1:8443", "path/to/server-key.pem", "path/to/server-cert.pem")
	srv.Run()
}
```