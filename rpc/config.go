package rpc

import (
	rpcSpec "commentService/rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Comment string `json:"comment" config:"comment,default=127.0.0.1:10001"`
}

func client(url string) (*grpc.ClientConn, error) {
	return grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func NewCommentService(c *Config) (rpcSpec.CommentServiceClient, error) {
	conn, err := client(c.Comment)
	if err != nil {
		return nil, err
	}
	return rpcSpec.NewCommentServiceClient(conn), nil
}
