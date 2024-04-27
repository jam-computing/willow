package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {

    packet := GoodPacket();
    packet.command = 11
    packet = packet.SendRecv();

    var a Animation = MakeAnimation(packet.data)

	component := hello(a.Name)
	http.Handle("/", templ.Handler(component))

    fmt.Println("Listening on http://localhost:3000")

	http.ListenAndServe(":3000", nil)

	// component.Render(context.Background(), os.Stdout);
}
