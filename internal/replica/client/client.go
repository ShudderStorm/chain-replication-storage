package client

import (
	"context"

	"github.com/ShudderStorm/chain-replication-storage/internal/grpc/replica/pb"
	"google.golang.org/grpc"
)

type Client struct{}

func (c Client) Store(ctx context.Context, in *pb.StoreRequest, opts ...grpc.CallOption) (*pb.StoreResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c Client) Load(ctx context.Context, in *pb.LoadRequest, opts ...grpc.CallOption) (*pb.LoadResponse, error) {
	//TODO implement me
	panic("implement me")
}
