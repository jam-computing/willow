package tcp

type Packet struct {
    Version int
}

func NewPacket() Packet {
    return Packet{}
}
