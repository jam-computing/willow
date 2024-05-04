package main

import (
	"encoding/json"

	"github.com/charmbracelet/log"
	"github.com/jam-computing/willow/pkg"
	"github.com/jam-computing/willow/pkg/player"
	"github.com/jam-computing/willow/pkg/protocol"
	"github.com/jam-computing/willow/pkg/webstate"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = pkg.NewTemplate()

	e.GET("/", func(c echo.Context) error {
		packet := protocol.NewFullPacket(protocol.NewMetaPacket(), nil, nil)
		packet.Meta.Status = 200
		packet.Meta.Command = 4
		recv := packet.SendRecv()

		jsonData := recv.Data.Data[:recv.Meta.Len]

		var animations []player.Animation
		err := json.Unmarshal([]byte(jsonData), &animations)

		if err != nil {
			log.Error("could not unmarsharl the json data\n", "err", err)
			return c.Render(200, "index", recv.Data.Data)
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

		p := protocol.NewFullPacket(protocol.NewMetaPacket(), &title, nil)
		p.Meta.Status = 200
		p.Meta.Command = 2

		_ = p.SendRecv()

		if p.Meta.Status == 200 {
			playbarData := webstate.NewPlaybarData(title, artist)
			return c.Render(200, "player", playbarData)
		}

		return nil
	})

	log.Info("Listening on http://localhost:3000")

	e.File("/card.css", "../../views/css/card.css")
	e.File("/materialize.css", "../../views/css/materialize.css")
	e.File("/materialize.js", "../../views/js/materialize.js")
	e.Logger.Fatal(e.Start(":3000"))
}
