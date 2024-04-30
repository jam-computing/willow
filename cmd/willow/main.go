package main

import (
	"encoding/json"
	"fmt"

	"github.com/jam-computing/willow/pkg"
	"github.com/jam-computing/willow/pkg/webstate"
	"github.com/jam-computing/willow/pkg/player"
	"github.com/jam-computing/willow/pkg/protocol"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = pkg.NewTemplate()

	data := webstate.NewData()

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

		cardData := webstate.NewData()

		for _, a := range animations {
			cardData.Cards = append(cardData.Cards, webstate.NewCard(a.Title, a.Artist))
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
			playbarData := webstate.NewPlaybarData(title, artist)
			return c.Render(200, "player", playbarData)
		}

		return nil
	})

	fmt.Println("Listening on http://localhost:3000")

	e.File("/card.css", "../../views/css/card.css")
	e.File("/materialize.css", "../../views/css/materialize.css")
	e.File("/materialize.js", "../../views/js/materialize.js")
	e.Logger.Fatal(e.Start(":3000"))
}
