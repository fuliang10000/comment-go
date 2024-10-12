package rpcService

import (
	"commentService/internal/comment/dao/query"
	rpcSpec "commentService/rpc/proto"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommentService struct {
	rpcSpec.UnimplementedCommentServiceServer
	*query.Query
}

func (s CommentService) CreateComment(ctx context.Context, request *rpcSpec.CreateCommentRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func NewCommentService(query *query.Query) *CommentService {
	return &CommentService{
		UnimplementedCommentServiceServer: rpcSpec.UnimplementedCommentServiceServer{},
		Query:                             query,
	}
}
