package comment

import (
	rpcService "commentService/internal/comment/service/rpc"
	rpcSpec "commentService/rpc/proto"
	rpcServer "commentService/rpc/server"
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
