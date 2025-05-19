package main

import (
	"context"
)

func communicate(telnet_client TelnetClient, cancel context.CancelFunc) {
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
	}
}

func run_client(telnet_client TelnetClient) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go communicate(telnet_client, cancel)
	return ctx
}
