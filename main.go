package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func setupRoutes() {
	r := gin.Default()
	r.GET("/", homePage)
	r.GET("/ws", wsEndpoint)
	r.Run(":2004")
}

func homePage(ctx *gin.Context) {
	fmt.Fprintf(ctx.Writer, "Home page")
}

func wsEndpoint(ctx *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
	}

	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}
	log.Println("Connected")
	reader(ws)

}

func reader(conn *websocket.Conn) {
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(msg))

		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	setupRoutes()
}
