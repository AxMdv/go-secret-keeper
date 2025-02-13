package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"secret-keeper/internal/config"
	"secret-keeper/internal/handlers"
	"secret-keeper/internal/logger"
	"secret-keeper/internal/service/auth"
	"secret-keeper/internal/service/keeper"
	"secret-keeper/internal/storage/postgres"
	"syscall"

	"google.golang.org/grpc"
)

type App struct {
	GRPCServer *grpc.Server
	Config     *config.Config
	Storage    *postgres.PostgresStorage
}

func New(cfg config.Config) (*App, error) {

	err := logger.Init(cfg.LoggingLevel)
	if err != nil {
		return nil, err
	}
	pgStorage, err := postgres.NewPostgresStorage(cfg.DataBaseDSN)
	if err != nil {
		return nil, err
	}
	keeperService := keeper.NewKeeperService(pgStorage)
	authService := auth.NewAuthService(cfg.TokenTTLMinutes, cfg.JWTSecretKey)
	grpcHandler := handlers.NewGRPCKeeperHandler(keeperService, authService)
	grpcServer := handlers.NewRegisteredServer(grpcHandler)
	return &App{
		GRPCServer: grpcServer,
		Config:     &cfg,
		Storage:    pgStorage,
	}, nil
}

func (a *App) Run() error {
	go func() {
		if err := a.runGRPCServer(); err != grpc.ErrServerStopped {
			log.Fatal("error: in run server:", err)
		}
	}()
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	fmt.Println("waiting for ctrl+c")
	c := <-interruptChan
	fmt.Println("recieved signal: ", c)
	a.stopGRPCServer()
	a.Storage.Close()
	return nil
}

func (a *App) runGRPCServer() error {
	listen, err := net.Listen("tcp", a.Config.GRPCRunAddr)
	if err != nil {
		log.Fatalf("Failed to listen TCP %s: %s ", a.Config.GRPCRunAddr, err)
	}
	fmt.Println("Сервер gRPC начал работу")
	// получаем запрос gRPC
	if err := a.GRPCServer.Serve(listen); err != nil {
		log.Fatalf("Failed to run GRPC server %s: %s", a.Config.GRPCRunAddr, err)
	}
	return err
}

func (a *App) stopGRPCServer() {
	a.GRPCServer.GracefulStop()
}
