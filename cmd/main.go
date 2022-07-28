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
	"github.com/flavioltonon/gandalf/infrastructure/cryptography"
	"github.com/flavioltonon/gandalf/infrastructure/logger/zap"
	"github.com/flavioltonon/gandalf/infrastructure/presenter/json"
	"github.com/flavioltonon/gandalf/infrastructure/repository/mongo"
	"github.com/flavioltonon/gandalf/infrastructure/uuid"
	"github.com/gorilla/mux"
)

func main() {
	mongoClient, err := mongo.NewClient(context.Background(), "mongodb://database:27017")
	if err != nil {
		log.Fatal(err)
	}

	database := mongoClient.Database("gandalf")

	var (
		md5Encryptor             = cryptography.NewMD5Encryptor()
		uuidV4Factory            = uuid.NewV4Factory()
		presenter                = json.NewPresenter()
		logger, _                = zap.NewLogger()
		usersRepository          = database.NewUsersRepository()
		authenticationService    = services.NewAuthenticationService(usersRepository, uuidV4Factory, md5Encryptor)
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
