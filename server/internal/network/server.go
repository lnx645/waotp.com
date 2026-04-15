package network

import (
	"fmt"
	"time"

	"github.com/gorilla/mux"
	"github.com/zishang520/socket.io/servers/socket/v3"
)

const SOCKET_PATH string = "/realtime/ws/"

func InitServer() *mux.Router {
	r := mux.NewRouter()

	io := InitSocketServer()
	io.On("connection", func(clients ...any) {
		client := clients[0].(*socket.Socket)
		token, _ := client.Request().Query().Get("token")
		fmt.Println(token)
		fmt.Println("Client connected:", client.Id())
		client.Join(socket.Room("anonim"))
		time.AfterFunc(500*time.Millisecond, func() {
			client.Emit("status", "Connected to server websocket")
		})
	})
	r.Handle(SOCKET_PATH, io.ServeHandler(nil))
	return r
}
