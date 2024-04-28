package main

import (
	"fmt"

	"github.com/jam-computing/willow/pkg/tcp"
)

func main() {
    packet := tcp.NewPacket()
    packet.command = 11
    packet = packet.SendRecv();

}
