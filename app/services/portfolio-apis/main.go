package main

import (
	"context"
	"fmt"
	"html"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"sr-codefreak/portfolio-visualiser/foundation/logger"
)

func main() {

	log := logger.NewLogger()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// time.Sleep(time.Second * 1)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Println("Port Exposed")
		log.WithField("hey", "hey").Info("testing log")
		// time.Sleep(time.Second * 15)
	})

	s := &http.Server{
		Addr:           ":4567",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	servererr := make(chan error)
	go func() {
		fmt.Println("Server Started on port ", s.Addr)
		log.WithField("port", s.Addr).Info("server started")
		servererr <- s.ListenAndServe()
	}()

	select {
	case err := <-servererr:
		fmt.Println("server err, ", err)
	case sign := <-shutdown:
		fmt.Println("starting shutdown, ", sign)
		log.WithField("signal", sign).Info("shutdown started")
		defer fmt.Println("shutdown complete, ", sign)
		defer log.WithField("signal", sign).Info("shutdown completed")

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Asking listener to shutdown and shed load.
		if err := s.Shutdown(ctx); err != nil {
			s.Close()
			fmt.Println("could not stop server gracefully: ", err)
		}
	}

}

func shutdown() {
	panic("unimplemented")
}

func startService() {
	panic("unimplemented")
}

func setConfigs() {
	panic("unimplemented")
}
