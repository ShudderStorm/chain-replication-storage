package server

import (
	"context"

	"github.com/ShudderStorm/chain-replication-storage/internal/grpc/replica/pb"
)

type Server struct {
	pb.UnimplementedReplicaServer
}

func (s Server) Store(ctx context.Context, request *pb.StoreRequest) (*pb.StoreResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) Load(ctx context.Context, request *pb.LoadRequest) (*pb.LoadResponse, error) {
	//TODO implement me
	panic("implement me")
}
