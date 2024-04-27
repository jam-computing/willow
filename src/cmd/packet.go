package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

var PROTOCOL_VERSION byte = 1

type Packet struct {
	version byte
	command byte
	status  uint16
	id      uint16
	data    string
}

func GoodPacket() Packet {
	packet := Packet{}

	packet.version = PROTOCOL_VERSION
	packet.command = 1
	packet.status = 200
	packet.id = 0
	packet.data = ""

	return packet
}

func BadPacket() Packet {
	packet := Packet{}

	packet.version = PROTOCOL_VERSION
	packet.command = 0
	packet.status = 500
	packet.id = 0
	packet.data = ""

	return packet
}

func Make(bytes []byte) Packet {
	packet := GoodPacket()

	packet.version = bytes[0]
	packet.command = bytes[1]

	packet.status = binary.LittleEndian.Uint16([]byte{bytes[2], bytes[3]})
	packet.id = binary.LittleEndian.Uint16([]byte{bytes[4], bytes[5]})
	length := binary.LittleEndian.Uint16([]byte{bytes[6], bytes[7]})

	if length > 0 {
		packet.data = string(bytes[8 : 8+length])
	} else {
		packet.data = ""
	}

	return packet
}

func (packet *Packet) ToBytes() []byte {

	// This feels bad
	bytes := make([]byte, 0, 1024)
	bytes = append(bytes, packet.version)
	bytes = append(bytes, packet.command)

	status := make([]byte, 2)
	id := make([]byte, 2)

	binary.LittleEndian.PutUint16(status, packet.status)
	binary.LittleEndian.PutUint16(id, packet.id)

	bytes = append(bytes, status...)
	bytes = append(bytes, id...)

	if len(packet.data) > 0 {
		data_len := uint16(len(packet.data))
		len := make([]byte, 2)
		binary.LittleEndian.PutUint16(len, data_len)

		bytes = append(bytes, len...)
		bytes = append(bytes, packet.data...)
	} else {
		len := make([]byte, 2)
		binary.LittleEndian.PutUint16(len, 0)
		bytes = append(bytes, len...)
	}

	return bytes
}

func (packet *Packet) SendRecv() Packet {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println(err)
		return BadPacket()
	}

	_, err = conn.Write(packet.ToBytes())

	if err != nil {
		fmt.Println(err)
		return BadPacket()
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)

	if err != nil {
		fmt.Println(err)
		return BadPacket()
	}

	conn.Close()
	return Make(buf)
}
