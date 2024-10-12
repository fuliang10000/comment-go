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
	addr := fmt.Sprintf(":%d", s.Config.ServerPort)
	slog.Info("GRPC server正在运行...：" + addr)
	l, err := net.Listen("tcp", addr)
	util.PanicError(err)
	util.PanicError(s.Server.Serve(l))
}
