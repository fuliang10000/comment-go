package rpcServer

import (
	"commentService/app/config"
	"commentService/pkg/util"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type Server struct {
	*grpc.Server
	*config.Config
}

func (s Server) Run() {
	addr := fmt.Sprintf(":%d", s.ServerPort)
	slog.Info("GRPC.Server正在运行..." + addr)
	l, err := net.Listen("tcp", addr)
	util.PanicError(err)
	util.PanicError(s.Server.Serve(l))
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Server: grpc.NewServer(),
		Config: cfg,
	}
}
