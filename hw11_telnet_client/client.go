package main

import (
	"context"
	"errors"
	"fmt"
)

func clientGorutine(clientFunc func() error) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func(clientFunc func() error) {
		for {
			err := clientFunc()
			if errors.Is(err, &ClientClosedError{}) {
				fmt.Println(err.Error())
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

// func send(telnecClient TelnetClient, sigintCtx context.Context) context.Context {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	go func(telnetClient TelnetClient) {
// 		for {
// 			err := telnetClient.Send()
// 			if err != nil {
// 				fmt.Printf("error send: %s", err.Error())
// 				cancel()
// 				return
// 			}
// 		}
// 	}(telnecClient)
// 	return ctx
// }

// func receive(telnetClient TelnetClient, sigintCtx context.Context) context.Context {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	go func(telnetClient TelnetClient) {
// 		for {
// 			err := telnetClient.Receive()
// 			if err != nil {
// 				fmt.Printf("error receive: %s", err.Error())
// 				cancel()
// 				return
// 			}
// 		}
// 	}(telnetClient)
// 	return ctx
// }

func run_client(telnetClient TelnetClient) (context.Context, context.Context) {
	// send_context := send(telnetClient, sigintCtx)
	// receive_context := receive(telnetClient, sigintCtx)
	// return send_context, receive_context
	sendContext := clientGorutine(telnetClient.Send)
	receiveContext := clientGorutine(telnetClient.Receive)
	return sendContext, receiveContext
}
