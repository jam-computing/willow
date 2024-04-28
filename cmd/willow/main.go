package main

import "fmt"

import "github.com/jam-computing/willow/pkg/tcp"

func main() {
    p := tcp.NewPacket()

    fmt.Println("Hello")
    fmt.Println(p.Version)
}
