package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func run(addr string, timeout time.Duration) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-c
		fmt.Println("Получен сигнал SIGINT")
		cancel()
	}()

	telnetClient := NewTelnetClient(addr, timeout, os.Stdin, os.Stdout)
	if err := telnetClient.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sendCtx, receiveCtx := runClient(telnetClient)

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

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "таймаут")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Укажите адрес и порт")
		os.Exit(1)
	}
	addr := args[0] + ":" + args[1]
	run(addr, *timeout)
}
