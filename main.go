package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure that the response supports HTTP/2 streaming
	if r.ProtoMajor == 2 {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}
		// Set headers for the response
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		// Stream some data
		for i := 0; i < 10; i++ {
			// Send a chunk of data to the client
			fmt.Fprintf(w, "data: This is message %d\n\n", i+1)

			// Flush the buffer, sending the data immediately
			flusher.Flush()

			// Simulate some work by sleeping
			time.Sleep(1 * time.Second)
		}
	} else {
		fmt.Fprintf(w, "HTTP/2 not supported on this connection")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check the protocol version
	if r.ProtoMajor == 2 {
		fmt.Fprintf(w, "HTTP/2")
	} else if r.ProtoMajor == 1 && r.ProtoMinor == 1 {
		fmt.Fprintf(w, "HTTP/1.1")
	} else {
		fmt.Fprintf(w, "Unknown HTTP version: %s", r.Proto)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/stream", streamHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Starting server on :8080")
	if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
