package model

type User struct {
	Login    string
	Password string
}

type LoginPassword struct {
	ID        string
	UserLogin string
	Login     string
	Password  string
	Metadata  string
}

type TextData struct {
	ID        string
	UserLogin string
	Data      string
	Metadata  string
}

type BinaryData struct {
	ID        string
	UserLogin string
	Data      []byte
	Metadata  string
}

type BankCard struct {
	CardNumber int
	UserLogin  string
	Owner      string
	ExpDate    string
	CVV        int
	Metadata   string
}

type AllUserData struct {
	LoginPassword []LoginPassword
	TextData      []TextData
	BinaryData    []BinaryData
	BankCard      []BankCard
}
