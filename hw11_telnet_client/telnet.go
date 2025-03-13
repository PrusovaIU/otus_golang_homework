package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type TCPClient struct {
	conn    net.Conn
	in      io.ReadCloser
	out     io.Writer
	address string
	timeout time.Duration
}

func (c *TCPClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *TCPClient) Connect() error {
	if c.conn == nil {
		conn, err := net.DialTimeout("tcp", c.address, c.timeout)
		if err != nil {
			fmt.Println("Ошибка подключения:", err)
			return err
		}
		c.conn = conn
	} else {
		fmt.Println("Соединение уже установлено")
	}
	return nil
}

func (c *TCPClient) Send() error {
	if c.conn == nil {
		return fmt.Errorf("не было произведено подключение к серверу. Используйте функцию Connect()")
	}
	_, err := io.Copy(c.conn, c.in)
	return err
}

func (c *TCPClient) Receive() error {
	if c.conn == nil {
		return fmt.Errorf("не было произведено подключение к серверу. Используйте функцию Connect()")
	}
	_, err := io.Copy(c.out, c.conn)
	return err
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	client := &TCPClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
	return client
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.
