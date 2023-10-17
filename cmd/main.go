package main

import (
	"context"
	"doodocs/handler"
	"doodocs/service"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	
	// config, err := config.NewConfig()
	// if err != nil {
	// 	errLog.Fatal(err)
	// }
	// if err := config.SetVars(); err != nil {
	// 	errLog.Fatal(err)
	// }

	service := service.NewService()

	handler := handler.NewHandler(infoLog, errLog, service)

	server := handler.Server()

	shutdown := make(chan os.Signal, 1)
    signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

	go func() {
        sig := <-shutdown
        server.Shutdown()
		fmt.Println()
        infoLog.Println("Received signal:", sig)
		cancel()
    }()
	
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	go func() {
        if err := server.Listen(port); err != nil {
            errLog.Fatal(err)
        }
    }()

	<-ctx.Done()

	infoLog.Println("Shutting down server...")
}