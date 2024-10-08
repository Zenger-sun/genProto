package main

import (
	"fmt"
	"genProto/service"
	"log/slog"
	"os"

	"genProto/msg"
	"genProto/server"
)

const (
	PROTO_PATH  = "./msg/proto/msg.proto"
	SERVER_ADDR = "127.0.0.1:8080"

	DB_HOST     = "127.0.0.1"
	DB_PORT     = 3308
	DB_USERNAME = "root"
	DB_PASSWORD = "test111"
	DB_SCHEMA   = "data_character"
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
		conf := &server.Conf{
			Addr: SERVER_ADDR,
			DbConf: &server.DbConf{
				Host:     DB_HOST,
				Port:     DB_PORT,
				Username: DB_USERNAME,
				Password: DB_PASSWORD,
				Schema:   DB_SCHEMA,
			},
		}

		ctx := server.NewContext(conf)
		svc := service.NewSvc(ctx)
		svc.Init()

		ctx.Server()
		ctx.Setup()

	case "help":
		fmt.Println("cmd content:")
		fmt.Println("\tgo run main.go msg\t ----make factory.go from proto to struct.")
		fmt.Println("\tgo run main.go server\t ----start a tcp server using proto.")
	}
}
