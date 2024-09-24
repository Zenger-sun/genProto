package server

import (
	"log/slog"
	"net"
	"sync"

	"genProto/msg/pb"

	"google.golang.org/protobuf/proto"
)

type Handler func(msg proto.Message, conn net.Conn) error

type Router struct {
	mu      sync.RWMutex
	Handler map[pb.MsgType]Handler
}

func Server(addr string, router *Router) error {
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

func readTCP(conn net.Conn, router *Router) {
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

		router.mu.RLock()
		handler, ok := router.Handler[pb.MsgType(head.MsgType)]
		router.mu.RUnlock()
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

func Response(conn net.Conn, res pb.MsgType, msg proto.Message) {
	h := &Head{
		Len:     0,
		MsgType: uint16(res),
	}

	conn.Write(PackMsg(h, msg))
}
