package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-c
		fmt.Println("Получен сигнал SIGINT")
		cancel()
	}()

	telnetClient := NewTelnetClient("localhost:4242", 10*time.Second, os.Stdin, os.Stdout)
	if err := telnetClient.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sendCtx, receiveCtx := run_client(telnetClient)

	select {
	case <-sendCtx.Done():
		return
	case <-receiveCtx.Done():
		return
	case <-ctx.Done():
		fmt.Println("Выполнение программы завершено")
	}
	telnetClient.Close()
}
