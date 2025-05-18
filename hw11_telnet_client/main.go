package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func close_telnet_client(client TelnetClient, msg string) {
	fmt.Println(msg)
	client.Close()
}

func send_receive(telnet_client TelnetClient) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("Старт работы клиента")
		for {
			err := telnet_client.Send()
			if err != nil {
				close_telnet_client(telnet_client, err.Error())
				cancel()
				return
			}
			err = telnet_client.Receive()
			if err != nil {
				close_telnet_client(telnet_client, err.Error())
				cancel()
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()
	return ctx
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-c
		fmt.Println("Получен сигнал SIGINT")
		cancel()
	}()

	telnet_client := NewTelnetClient("localhost:4242", 10*time.Second, os.Stdin, os.Stdout)
	if err := telnet_client.Connect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err_ctx := send_receive(telnet_client)

	select {
	case <-err_ctx.Done():
		return
	case <-ctx.Done():
		close_telnet_client(telnet_client, "Выполнение программы завершено")
	}
}
