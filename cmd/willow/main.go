package main

import (
	"fmt"

	"github.com/jam-computing/willow/pkg"
	"github.com/jam-computing/willow/pkg/protocol"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Card struct {
	Title  string
	Artist string
}

func newCard(name, artist string) Card {
	return Card{
		Title:  name,
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

type ErrorData struct {
	Error string
}

func newErrorData(e string) ErrorData {
	return ErrorData{
		Error: e,
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = pkg.NewTemplate()

	data := newData()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.POST("/play", func(c echo.Context) error {
		title := c.FormValue("title")

		p := protocol.NewPacket()
		p.Status = 200
		p.Command = 2
		p.Data = title
		recv := p.SendRecv()

        fmt.Printf("Status: %s\n", recv.Data)

		if len(recv.Data) > 0 {
            errorData := newErrorData(recv.Data)
			return c.Render(200, "error", errorData)
		}

		return nil
	})

	fmt.Println("Listening on http://localhost:3000")

	e.Logger.Fatal(e.Start(":3000"))
}
