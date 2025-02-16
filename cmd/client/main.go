package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"secret-keeper/internal/config"
	pb "secret-keeper/internal/proto"
	"secret-keeper/pkg/build"

	"secret-keeper/pkg/client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	build.PrintBuildInfo()
	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}
	// устанавливаем соединение с сервером
	conn, err := grpc.NewClient(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// получаем переменную интерфейсного типа UsersClient,
	// через которую будем отправлять сообщения

	// регаем прерывание
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	fmt.Println("waiting for ctrl+c")

	// запуск лупа
	grpcClient := pb.NewSecretKeeperClient(conn)
	cli := &client.Client{
		GRPCClient: grpcClient,
		Authed:     false,
		Config:     cfg,
	}
	go cli.InfiniteProcess()

	// полуичили прерывание
	c := <-interruptChan
	fmt.Println("recieved signal: ", c)
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
