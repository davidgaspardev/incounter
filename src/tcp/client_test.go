package tcp

import (
	"testing"

	log "github.com/davidgaspardev/golog"
)

func Test_Listen(t *testing.T) {
	tcpClient, err := CreateTcpClient("localhost", 8888)
	if err != nil {
		t.Error(err)
	}

	err = tcpClient.Send([]byte("/register cf23df2207d99a74fbe169e3eba035e633b65d94"))
	if err != nil {
		t.Error(err)
	}

	tcpClient.Listen(func(data []byte) {
		log.Info("TCP Client", "data received: "+string(data))
	})
}
