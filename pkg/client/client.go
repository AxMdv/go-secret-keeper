package client

import (
	"bufio"
	"fmt"
	"os"
	"secret-keeper/internal/config"
	pb "secret-keeper/internal/proto"
	"time"
)

const (
	// STATES OF CLIENT:
	NOT_AUTHED      = 0
	AUTHED          = 1
	SHOW_LOGIN_PASS = 2
	SHOW_TEXT_DATA  = 3
	SHOW_BIN_DATA   = 4
	SHOW_BANK_CARDS = 5
	EXIT            = 1337
)

type Client struct {
	Registered bool
	Authed     bool
	GRPCClient pb.SecretKeeperClient
	WantToExit bool
	Config     *config.Config
	JWTToken   string
}

func (c *Client) InfiniteProcess() {
	fmt.Println("----------Сервис secret keeper------------")
	time.Sleep(250 * time.Millisecond)
	state := NOT_AUTHED
	for {
		switch state {
		case NOT_AUTHED:
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("1 - регистрация")
			time.Sleep(250 * time.Millisecond)
			fmt.Println("2 - войти в систему")
			time.Sleep(250 * time.Millisecond)
			fmt.Println("000 - завершить работу")
			time.Sleep(250 * time.Millisecond)
			fmt.Print("Введите значение: ")
			scanner.Scan()
			text := scanner.Text()
			switch text {
			case "1":
				fmt.Println("Вы выбрали регистрацию")
			case "2":
				fmt.Println("Вы выбрали войти в систему")
			case "000":
				fmt.Println("Вы выбрали выйти из приложения, продолжить?")
				fmt.Println("1 - Да")
				fmt.Println("2 - Выбрать другой варант")
				state = EXIT
			default:
				fmt.Println("Введено неверное значение, попробуйте ещё раз!")
				time.Sleep(1 * time.Second)
				continue
			}
			// Это главное меню типа
		case AUTHED:
			fmt.Println("----------Сервис secret keeper------------")
			time.Sleep(250 * time.Millisecond)
			fmt.Println("----------Главное меню------------")
			time.Sleep(250 * time.Millisecond)
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Println("1 - показать логины и пароли")
			time.Sleep(250 * time.Millisecond)

			fmt.Println("2 - показать текстовые данные")
			time.Sleep(250 * time.Millisecond)

			fmt.Println("3 - показать бинарные данные")
			time.Sleep(250 * time.Millisecond)

			fmt.Println("4 - показать данные банковских карт")
			time.Sleep(250 * time.Millisecond)

			fmt.Println("000 - завершить работу")
			time.Sleep(250 * time.Millisecond)

			fmt.Print("Введите значение: ")
			scanner.Scan()
			text := scanner.Text()
			switch text {
			case "1":
				fmt.Println("Вы выбрали показать логины и пароли")
				state = SHOW_LOGIN_PASS
				// c.GRPCClient.GetUserStoredData(context.Background(), )
			case "2":
				fmt.Println("Вы выбрали показать текстовые данные")
				state = SHOW_TEXT_DATA
			case "3":
				fmt.Println("Вы выбрали показать бинарные данные")
				state = SHOW_BIN_DATA
			case "4":
				fmt.Println("Вы выбрали показать данные банковских карт")
				state = SHOW_BANK_CARDS
			case "000":
				fmt.Println("Вы выбрали выйти из приложения, продолжить?")
				state = EXIT
			default:
				fmt.Println("Введено неверное значение, попробуйте ещё раз!")
				time.Sleep(1 * time.Second)
				continue
			}
		case SHOW_BANK_CARDS:
			// тут берем данные с сервера / запрашиваем данные в другой рутине и здесь мы вообще не ходим на сервак
		case SHOW_TEXT_DATA:
		case SHOW_BIN_DATA:
		case SHOW_BANK_CARDS:
		}

	}
}

func scan() {
	// To create dynamic array
	arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Text: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
			arr = append(arr, text)
		} else {
			break
		}

	}
	// Use collected inputs
	fmt.Println(arr)
}
