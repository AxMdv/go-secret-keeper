package main

import (
	"log"
	"secret-keeper/internal/app"
	"secret-keeper/internal/config"
	"secret-keeper/pkg/build"
)

// go run -ldflags "-X main.buildVersion=v1.0.1 -X 'main.buildDate=$(date +'%Y/%m/%d %H:%M:%S')' -X main.buildCommit=hello world" main.go
// go run -ldflags "-X main.buildVersion=v1.0.1 -X main.buildCommit=hello-world" main.go -f=""

func main() {
	build.PrintBuildInfo()
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("failed to parse config app: %s", err.Error())
	}
	a, err := app.New(*cfg)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
	log.Println("app exited properly")
}
