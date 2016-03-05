package handlers

import(
	"golang.org/x/net/websocket"
	"io"
	"fmt"
)

func EchoServer(ws *websocket.Conn) {
	fmt.Printf("got websocket connection %#v\n", *ws)
	io.Copy(ws, ws)
}