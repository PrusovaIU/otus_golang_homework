package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClientGorutine(t *testing.T) {
	cases := []struct {
		name       string
		clientFunc func() error
	}{
		{
			name: "client closed",
			clientFunc: func() error {
				return &ClientClosedError{Message: "Соединение закрыто клиентом"}
			},
		},
		{
			name: "client error",
			clientFunc: func() error {
				return fmt.Errorf("Соединение закрыто сервером")
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := clientGorutine(tc.clientFunc)
			childCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
			defer cancel()
			<-childCtx.Done()
			require.ErrorIs(t, childCtx.Err(), context.Canceled)
		})
	}

	// t.Run("client closed error case", func(t *testing.T) {
	// 	clientFunc := func() error {
	// 		return &ClientClosedError{Message: "Соединение закрыто клиентом"}
	// 	}
	// 	ctx := clientGorutine(clientFunc)
	// 	childCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// 	defer cancel()
	// 	<-childCtx.Done()
	// 	require.ErrorAs(t, childCtx.Err(), &context.Canceled)
	// })

}
