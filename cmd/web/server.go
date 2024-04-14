package web

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func InitServer(handler http.Handler) {

	srv := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		WriteTimeout: time.Duration(15) * time.Second,
		ReadTimeout:  time.Duration(15) * time.Second,
		IdleTimeout:  time.Duration(15) * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Wait for signal to shut down server
	shutDownChannel := make(chan os.Signal, 1)
	signal.Notify(shutDownChannel, os.Interrupt)
	<-shutDownChannel
}
