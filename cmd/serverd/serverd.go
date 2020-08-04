package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(time.Duration(50+rand.Intn(250)) * time.Millisecond) // Artificial Latency [50-300)ms
		w.WriteHeader(http.StatusOK)
	})

	server := http.Server{
		Addr:    "localhost:9000",
		Handler: mux,
	}

	serverErrors := make(chan error)
	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	select {
	case e := <-serverErrors:
		log.Printf("server error: %v\n", e)
		return

	case <-osSignals:
		signal.Reset(os.Interrupt, syscall.SIGTERM)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("unable to gracefully shut down server: %v\n", err)

		log.Println("attempting to forcefully shut down server")
		if err := server.Close(); err != nil {
			log.Printf("error forcefully shutting down server: %v\n", err)
		}
	}
}
