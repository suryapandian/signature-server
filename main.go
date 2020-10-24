package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"signatureserver/config"
	"signatureserver/handlers"
	"syscall"
)

func main() {
	server := &http.Server{
		Addr:    ":" + config.PORT,
		Handler: handlers.GetRouter(),
	}

	go func(server *http.Server) {
		log.Printf("server running on: %s", config.PORT)
		if err := server.ListenAndServe(); err != nil {
			log.Printf("server listen error: %s", err)
		}
	}(server)

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-stopCh
	log.Printf("gracefully shutting down server")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("error shutting server down gracefully")
	}
}
