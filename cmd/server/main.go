package main

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"go-labs-game-platform/internal/bootstrap"
	"go-labs-game-platform/internal/config"
	"go-labs-game-platform/internal/httpserver"
)

func main() {
	_, err := config.New()
	if err != nil {
		panic(err)
	}

	deps, err := bootstrap.Up()
	if err != nil {
		panic(err)
	}

	s := httpserver.New(deps)

	server := &http.Server{
		Addr:           s.Addr(),
		Handler:        s.Router(),
		ReadTimeout:    config.Get().HTTP.ReadTimeout,
		WriteTimeout:   config.Get().HTTP.WriteTimeout,
		IdleTimeout:    time.Second * 10,
		MaxHeaderBytes: 256,
	}

	logrus.Info("Server http listening")

	panic(server.ListenAndServe())
}
