package comment

import (
	"commentService/app/config"
	rpcService "commentService/internal/comment/service/rpc"
	"commentService/pkg/util"
	"commentService/rpc"
	rpcSpec "commentService/rpc/proto"
	rpcServer "commentService/rpc/server"
	"go.uber.org/dig"
	"log/slog"
)

type Server struct {
	*rpcServer.Server
	*rpcService.CommentService
}

func (s Server) Run() {
	slog.Info("comment.Server.Run")
	rpcSpec.RegisterCommentServiceServer(s.Server, s.CommentService)
	s.Server.Run()
}

func NewServer(rpc *rpcServer.Server, comment *rpcService.CommentService) *Server {
	return &Server{
		Server:         rpc,
		CommentService: comment,
	}
}

func Run() {
	container := dig.New()

	util.PanicError(container.Provide(config.NewConfig))
	util.PanicError(container.Provide(config.RpcConfig))
	util.PanicError(container.Provide(QueryInit))
	util.PanicError(container.Provide(rpc.NewCommentService))
	util.PanicError(container.Provide(rpcService.NewCommentService))
	util.PanicError(container.Provide(NewServer))
	util.PanicError(container.Provide(rpcServer.NewServer))

	util.PanicError(container.Invoke(func(server *Server) {
		server.Run()
	}))
}
