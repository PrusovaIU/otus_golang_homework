package main

import (
	"context"
	"fmt"
)

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

func run_client(telnetClient TelnetClient) (context.Context, context.Context) {
	sendContext := clientGorutine(telnetClient.Send)
	receiveContext := clientGorutine(telnetClient.Receive)
	return sendContext, receiveContext
}
