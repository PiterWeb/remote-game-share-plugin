package main

import "C"
import (
	"github.com/PiterWeb/remote-controller-game-share-plugin/src"
)

//export init_host
func init_host(app_port uint, port uint, protocol uint) {
	switch protocol {
	case src.UDP:
		src.UdpLogic(app_port, port)
	case src.TCP:
		src.TcpLogic(app_port, port)
	default:
		return
	}
}

//export init_client
func init_client(app_port uint, port uint, protocol uint) {
	// init_host has the same logic as init_client would have
	init_host(app_port, port, protocol)
}

//export background
func background() {}

func main() {}
