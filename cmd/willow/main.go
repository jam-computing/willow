package main

import (
	"fmt"

	"github.com/jam-computing/willow/pkg/protocol"
	"github.com/jam-computing/willow/pkg/player"
	"github.com/jam-computing/willow/pkg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Data struct {
    Title string
}

func main() {
    packet := protocol.NewPacket();
    packet.Command = 11
    recv := packet.SendRecv()

    animation := player.MakeAnimation(recv.Data)

    e := echo.New();
    e.Use(middleware.Logger())

    e.Renderer = pkg.NewTemplate()

    data := Data{ Title: animation.Name }

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", data)
    })

    fmt.Println("Listening on http://localhost:3000")

    e.Logger.Fatal(e.Start(":3000"))
}
