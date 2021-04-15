package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/http/rest"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/listing"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/storage/mongo"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/storage/redis"
)

var (
	// defaultHost  = "0.0.0.0:8081"
	defaultHost  = "localhost:8081"
	inProduction = false
)

func main() {

	// Initialize storages
	lr := new(mongo.ListRepo)
	lc := new(redis.ListCache)
	lister := listing.NewService(lr, lc)

	// Initialize router
	router := rest.Handler(lister)

	host := os.Getenv("API_HOST")
	if host == "" {
		host = defaultHost
	}

	srv := &http.Server{
		Addr:         host,
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	// when testing locally it doesn't make sense to start
	// HTTPS server, so only do it in production.
	// In real code, I control this with -production cmd-line flag
	// if inProduction {
	// 	m := autocert.Manager{
	// 		Prompt:     autocert.AcceptTOS,
	// 		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
	// 		Cache:      autocert.DirCache("/var/www/.cache"),
	// 	}

	// 	go func() {
	// 		log.Fatal(autotls.RunWithManager(router, &m))
	// 	}()
	// } else {

	// }

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
