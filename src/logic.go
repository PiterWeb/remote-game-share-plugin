package src

import (
	"fmt"
	"net"

	"github.com/nats-io/nats.go"
)

func TcpLogic(app_port uint, port uint) {
	globalLogic(TCP, app_port, port)
}

func globalLogic(protocol, app_port uint, port uint) {

	addr := fmt.Sprintf("127.0.0.1:%d", port)

	var conn net.Conn
	var err error

	switch protocol {
	case UDP:
		conn, err = net.Dial("udp", addr)
	case TCP:
		conn, err = net.Dial("tcp", addr)
	default:
		conn, err = nil, fmt.Errorf("invalid protocol")
	}

	if err != nil {
		return
	}

	defer conn.Close()

	natsUrl := fmt.Sprintf("nats://127.0.0.1:%d", app_port)

	nc, err := nats.Connect(natsUrl)

	if err != nil {
		return
	}

	defer nc.Drain()

	nc.Subscribe(INPUT_SUBJECT, func(msg *nats.Msg) {
		conn.Write(msg.Data)
	})

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			break
		}

		nc.Publish(OUTPUT_SUBJECT, buf[:n])
	}
}

func UdpLogic(app_port uint, port uint) {
	globalLogic(UDP, app_port, port)
}
