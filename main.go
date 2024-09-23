package main

import (
	"fmt"
	"log/slog"
	"os"

	"genProto/msg"
	"genProto/msg/pb"
	"genProto/server"
	"genProto/service"
)

const (
	PROTO_PATH  = "./msg/proto/msg.proto"
	SERVER_ADDR = "127.0.0.1:8080"
)

func main() {
	if len(os.Args) <= 1 {
		slog.Warn("cmd not set! use 'go run main.go help' get other cmd.")
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "msg":
		msg.GenMsg(PROTO_PATH)

	case "server":
		router := make(map[pb.MsgType]server.Handler)
		router[pb.MsgType_MSG_LOGIN_REQ] = service.Login

		err := server.Server(SERVER_ADDR, router)
		if err != nil {
			slog.Warn("tcp server runtime err: ", err)
			os.Exit(1)
		}

	case "help":
		fmt.Println("cmd content:")
		fmt.Println("\tgo run main.go msg\t ----make factory.go from proto to struct.")
		fmt.Println("\tgo run main.go server\t ----start a tcp server using proto.")
	}
}
