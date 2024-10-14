package rpcService

import (
	"commentService/internal/comment/repo"
	rpcSpec "commentService/rpc/proto"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommentService struct {
	rpcSpec.UnimplementedCommentServiceServer
	*repo.Comment
}

func (s CommentService) CreateComment(ctx context.Context, request *rpcSpec.CreateCommentRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func NewCommentService(comment *repo.Comment) *CommentService {
	return &CommentService{
		UnimplementedCommentServiceServer: rpcSpec.UnimplementedCommentServiceServer{},
		Comment:                           comment,
	}
}
