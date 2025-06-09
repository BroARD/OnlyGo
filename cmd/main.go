package main

import (
	"OnlyGo/logging"
	"OnlyGo/pkg/config"
	"OnlyGo/pkg/db"
	"OnlyGo/pkg/quote"
	"OnlyGo/pkg/user"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

const (
	SocketFile = "app.sock"
	Sock = "sock"
)

func main() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()

	logger.Info("Start db")
	database, err := db.InitDB()
	if err != nil {
		logger.Info("Could not start a DB")
		logger.Fatal()
	}

	logger.Info("Start MongoDB")
	mongoDB, err := db.InitMongoDB()
	if err != nil {
		logger.Info("Could not start a MongoDB")
		logger.Fatal()
	}

	userRepository := user.NewRepository(mongoDB)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)


	logger.Info("registring a quote repository")
	quoteRepository := quote.NewRepository(database)
	logger.Info("registring a quote service")
	quoteService := quote.NewService(quoteRepository)
	logger.Info("registring a quote handlers")
	quoteHandler := quote.NewHandler(quoteService, logger)

	r := mux.NewRouter()
	quoteHandler.Register(r)
	userHandler.Register(r)

	start(r, cfg)
}

func start(router *mux.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	
	var listener net.Listener
	var listenError error

	if cfg.Listen.Type == Sock {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("Create socket")
		socketPath := path.Join(appDir, SocketFile)
		logger.Debugf("socket path: %s", socketPath)

		listener, listenError = net.Listen("unix", socketPath)	
	} else {
		listener, listenError = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	}

	if listenError != nil {
		logger.Fatal(listenError)
	}

	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
