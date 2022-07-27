package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/flavioltonon/gandalf/adapters/http/controllers"
	"github.com/flavioltonon/gandalf/application/services"
	"github.com/flavioltonon/gandalf/infrastructure/logger/zap"
	"github.com/flavioltonon/gandalf/infrastructure/presenter/json"
	"github.com/flavioltonon/gandalf/infrastructure/repository/memory"
	"github.com/gorilla/mux"
)

func main() {
	var (
		presenter                = json.NewPresenter()
		logger, _                = zap.NewLogger()
		usersRepository          = memory.NewUsersRepository()
		authenticationService    = services.NewAuthenticationService(usersRepository)
		authenticationController = controllers.NewAuthenticationController(authenticationService, presenter, logger)
	)

	router := mux.NewRouter()
	router.HandleFunc("/register", authenticationController.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/login", authenticationController.Login).Methods(http.MethodPost)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Printf("Listening and serving at %s\n", server.Addr)
		server.ListenAndServe()
	}()

	interrupt := make(chan os.Signal, 1)

	signal.Notify(
		interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-interrupt

	log.Println("Gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
