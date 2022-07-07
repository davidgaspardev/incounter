package tcp

import (
	"bufio"
	"fmt"
	"net"
)

type _Client struct {
	conn *net.TCPConn
}

func CreateTcpClient(addr string, port uint16) (*_Client, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}

	tcpClient := &_Client{
		conn: conn,
	}

	return tcpClient, nil
}

func (client *_Client) Send(data []byte) error {
	if data[len(data)-1] != '\n' {
		data = append(data, '\n')
	}
	_, err := client.conn.Write(data)

	return err
}

func (client *_Client) Listen(action func([]byte)) {
	go func() {
		for {
			data, err := bufio.NewReader(client.conn).ReadBytes('\n')
			if err != nil {
				continue
			}

			action(data[:len(data)-1])
		}
	}()
}
