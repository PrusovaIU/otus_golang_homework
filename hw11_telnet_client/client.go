package main

import (
	"context"
	"fmt"
)

// clientGorutine запускает функцию clientFunc в отдельной горутине и возвращает контекст с возможностью отмены.
func clientGorutine(clientFunc func() error) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func(clientFunc func() error) {
		for {
			err := clientFunc()
			if err, ok := err.(*ClientClosedError); ok {
				fmt.Println(err.Error())
				cancel()
				return
			}
			if err != nil {
				fmt.Printf("client error: %s", err.Error())
				cancel()
				return
			}
		}
	}(clientFunc)
	return ctx
}

// runClient запускает функции Send и Receive в отдельных горутинах и возвращает контексты с возможностью отмены.
func runClient(telnetClient TelnetClient) (context.Context, context.Context) {
	sendContext := clientGorutine(telnetClient.Send)
	receiveContext := clientGorutine(telnetClient.Receive)
	return sendContext, receiveContext
}
