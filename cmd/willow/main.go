package main

import (
	"fmt"

	"github.com/jam-computing/willow/pkg"
	"github.com/jam-computing/willow/pkg/player"
	"github.com/jam-computing/willow/pkg/protocol"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Card struct {
	Title  string
	Artist string
}

func newCard(name, artist string) Card {
    return Card {
        Title: name,
        Artist: artist,
    }
}

type Cards = []Card

type Data struct {
    Cards Cards
}

func newData() Data {
    return Data{
        Cards: []Card{
            newCard("Title One", "Author One"),
            newCard("Title Two", "Author Two"),
        },
    }
}

func main() {
	packet := protocol.NewPacket()
	packet.Command = 11
	recv := packet.SendRecv()

	_ = player.MakeAnimation(recv.Data)

	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = pkg.NewTemplate()

	data := newData()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

    e.POST("/play", func(c echo.Context) error {
        title := c.FormValue("title")
        artist := c.FormValue("artist")

        fmt.Println(title)
        fmt.Println(artist)

        p := protocol.NewPacket()
        p.Status = 200
        p.Command = 11
        recv := p.SendRecv()

        fmt.Println(recv.Status)

		return c.Render(200, "index", data)


	})

	fmt.Println("Listening on http://localhost:3000")

	e.Logger.Fatal(e.Start(":3000"))
}
