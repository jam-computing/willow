package main

import (
	"encoding/json"
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
	return Card{
		Title:  name,
		Artist: artist,
	}
}

type Cards = []Card

type Data struct {
	Cards Cards
    Title string
    Artist string
}

func newData() Data {
	return Data{
        Title: "",
        Artist: "",
    }
}

type PlaybarData struct {
	Title  string
	Artist string
}

func newPlaybarData(title, artist string) *PlaybarData {
    return &PlaybarData{
        Title: title,
        Artist: artist,
    }
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = pkg.NewTemplate()

	data := newData()

	e.GET("/", func(c echo.Context) error {
		p := protocol.NewPacket()
		p.Status = 200
		p.Command = 4
		recv := p.SendRecv()

		jsonData := recv.Data[:recv.Len]
		var animations []player.Animation
		err := json.Unmarshal([]byte(jsonData), &animations)

		if err != nil {
			fmt.Println(err)
			fmt.Printf("could not unmarsharl the json data\n")
			return c.Render(200, "index", data)
		}

		cardData := newData()

		for _, a := range animations {
			cardData.Cards = append(cardData.Cards, newCard(a.Title, a.Artist))
		}

		return c.Render(200, "index", cardData)
	})

	e.POST("/play", func(c echo.Context) error {
		title := c.FormValue("title")
		artist := c.FormValue("artist")

		p := protocol.NewPacket()
		p.Status = 200
		p.Command = 2
		p.Data = title
		recv := p.SendRecv()

		if len(recv.Data) > 0 {
			playbarData := newPlaybarData(title, artist)
			return c.Render(200, "player", playbarData)
		}

		return nil
	})

	fmt.Println("Listening on http://localhost:3000")

	e.File("/card.css", "../../css/card.css")
	e.File("/materialize.css", "../../css/materialize.css")
	e.File("/materialize.js", "../../js/materialize.js")
	e.Logger.Fatal(e.Start(":3000"))
}
