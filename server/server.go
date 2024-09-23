package server

import (
	"log/slog"
	"net"

	"genProto/msg/pb"

	"google.golang.org/protobuf/proto"
)

type Handler func(msg proto.Message, conn net.Conn) error

func Server(addr string, router map[pb.MsgType]Handler) error {
	laddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}

	l, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go readTCP(conn, router)
	}
}

func readTCP(conn net.Conn, router map[pb.MsgType]Handler) {
	buff := make([]byte, PACK_MAX_LEN)

	for {
		_, err := conn.Read(buff)
		switch err {
		case nil:
		default:
			slog.Warn("readTCP open conn err: ", err)
			return
		}

		head, msg, err := UnpackMsg(buff)
		if err != nil {
			slog.Warn("readTCP unpack head err: ", err)
			continue
		}

		handler, ok := router[pb.MsgType(head.MsgType)]
		if !ok {
			continue
		}

		err = handler(msg, conn)
		if err != nil {
			slog.Warn("handler err: ", err)
			continue
		}
	}
}

func Response(conn net.Conn, msg []byte) {
	conn.Write(msg)
}