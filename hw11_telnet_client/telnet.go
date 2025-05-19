package main

import (
	"bufio"
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
	conn         net.Conn
	conn_scanner bufio.Scanner
	in           bufio.Scanner
	out          bufio.Writer
	address      string
	timeout      time.Duration
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
		fmt.Println("Соединение установлено")
		c.conn = conn
		c.conn_scanner = *bufio.NewScanner(conn)
	} else {
		fmt.Println("Соединение уже установлено")
	}
	return nil
}

func (c *TCPClient) write(message []byte, to io.Writer) error {
	sent_bytes := 0
	message = append(message, []byte("\n")...)
	for sent_bytes < len(message) {
		n, err := to.Write(message[sent_bytes:])
		if err != nil {
			return err
		}
		sent_bytes += n
	}
	return nil
}

func (c *TCPClient) Send() error {
	if c.conn == nil {
		return fmt.Errorf("не было произведено подключение к серверу. Используйте функцию Connect()")
	}
	fmt.Println("Введите сообщение:")
	if c.in.Scan() {
		message := c.in.Bytes()
		return c.write(message, c.conn)
	}
	err := c.in.Err()
	return err
}

func (c *TCPClient) Receive() error {
	if c.conn == nil {
		return fmt.Errorf("не было произведено подключение к серверу. Используйте функцию Connect()")
	}
	if c.conn_scanner.Scan() {
		message := c.conn_scanner.Bytes()
		err := c.write(message, &c.out)
		if err != nil {
			return err
		}
		return c.out.Flush()
	}
	return c.conn_scanner.Err()
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	client := &TCPClient{
		address: address,
		timeout: timeout,
		in:      *bufio.NewScanner(in),
		out:     *bufio.NewWriter(out),
	}
	return client
}
