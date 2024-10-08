package server

import (
	"database/sql"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type Context struct {
	*sql.DB
	*Router
	*Conf
}

func (ctx *Context) Server() {
	db, err := DbServer(ctx.Conf.DbConf)
	if err != nil {
		slog.Error("db err: ", err)
	}

	ctx.DB = db

	go listen(ctx)
	slog.Info("server started!")
}

func (ctx *Context) Setup() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		switch <-exit {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			ctx.Shutdown()
			os.Exit(0)
		default:
			os.Exit(1)
		}
	}
}

func (ctx *Context) Shutdown() {
	slog.Info("Shutdown...")
	slog.Info("Closed!")
}

func NewContext(conf *Conf) *Context {
	sync := &Context{Conf: conf, Router: NewRouter()}
	return sync
}
