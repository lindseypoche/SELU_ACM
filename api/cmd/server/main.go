package main

import (
	"net/http"
	"time"

	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/http/rest"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/listing"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/storage/mongo"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/storage/redis"
)

var (
	port = "8081"
)

func main() {

	lr := new(mongo.ListRepo)
	lc := new(redis.ListCache)
	lister := listing.NewService(lr, lc)

	router := rest.Handler(lister)
	srv := &http.Server{
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
