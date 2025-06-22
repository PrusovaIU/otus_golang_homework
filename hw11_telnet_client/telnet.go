package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

type ClientClosedError struct {
	Message string
}

func (e *ClientClosedError) Error() string {
	return e.Message
}

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type TCPClient struct {
	conn       net.Conn
	connReader bufio.Reader
	in         bufio.Reader
	out        bufio.Writer
	address    string
	timeout    time.Duration
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
		c.connReader = *bufio.NewReader(conn)
	} else {
		fmt.Println("Соединение уже установлено")
	}
	return nil
}

func (c *TCPClient) write(message []byte, to io.Writer) error {
	sentBytes := 0
	for sentBytes < len(message) {
		n, err := to.Write(message[sentBytes:])
		if err != nil {
			return err
		}
		sentBytes += n
	}
	return nil
}

func (c *TCPClient) Send() error {
	if c.conn == nil {
		return fmt.Errorf("не было произведено подключение к серверу. Используйте функцию Connect()")
	}
	text, err := c.in.ReadString('\n')
	fmt.Printf("Send err: %v\n", err)
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		return &ClientClosedError{Message: "Соединение закрыто клиентом"}
	} else if err != nil {
		return err
	}
	err = c.write([]byte(text), c.conn)
	return err
}

func (c *TCPClient) Receive() error {
	if c.conn == nil {
		return fmt.Errorf("не было произведено подключение к серверу. Используйте функцию Connect()")
	}
	text, err := c.connReader.ReadString('\n')
	fmt.Printf("Receive err: %v\n", err)
	if err != nil {
		return err
	}
	err = c.write([]byte(text), &c.out)
	if err != nil {
		return err
	}
	err = c.out.Flush()
	return err
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	client := &TCPClient{
		address: address,
		timeout: timeout,
		in:      *bufio.NewReader(in),
		out:     *bufio.NewWriter(out),
	}
	return client
}
