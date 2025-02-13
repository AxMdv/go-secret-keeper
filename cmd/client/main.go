package main

import (
	"context"
	"log"

	pb "secret-keeper/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// устанавливаем соединение с сервером
	conn, err := grpc.NewClient(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// получаем переменную интерфейсного типа UsersClient,
	// через которую будем отправлять сообщения
	client := pb.NewSecretKeeperClient(conn)
	client.AuthUser(context.Background(), &pb.AuthUserRequest{})

	// функция, в которой будем отправлять сообщения
	// TestUsers(c)

}
