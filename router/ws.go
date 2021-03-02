package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"ice/app/ws"
	"net/http"
)

func InitWs(r *gin.RouterGroup)  {
	Router := r.Group("/api.v1/ws")
	{
		Router.GET("bigwall", func(c *gin.Context) {
			conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
			if error != nil {
				http.NotFound(c.Writer, c.Request)
				return
			}
			// websocket connect
			client := &ws.Client{ID: "11111", Socket: conn, Send: make(chan []byte)}
			ws.Manager.Register <- client
			go client.Read()
			go client.Write()
		})
	}
}

