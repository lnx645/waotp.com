package network

import (
	"fmt"
	"time"

	"github.com/zishang520/socket.io/servers/socket/v3"
	"github.com/zishang520/socket.io/v3/pkg/types"
)

var IO *socket.Server

func InitSocketServer() *socket.Server {
	config := socket.DefaultServerOptions()

	config.SetPingInterval(25 * time.Second)
	config.SetPingTimeout(10 * time.Second)

	config.SetMaxHttpBufferSize(1e6) // 1MB
	config.SetConnectTimeout(10 * time.Second)

	config.SetCors(&types.Cors{
		Origin:         "*",
		Credentials:    true,
		AllowedHeaders: []string{"Authorization"},
	})

	// Passing nil as the first argument uses the default creator
	IO = socket.NewServer(nil, config)

	return IO
}

func SendToRoom(room string, ev string, data interface{}) {
	if IO != nil {
		IO.To(socket.Room(room)).Emit(ev, data)
	} else {
		fmt.Println("Socket belum di definisikan")
	}
}
