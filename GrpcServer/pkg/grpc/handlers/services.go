package grpcHandlers

import (
	"context"
	"github.com/soheil-89/OMPFinex/GrpcServer/pkg/grpc/proto"
)

type Server struct {
	ChunkPb.UnimplementedChunkServer
}

func (s *Server) Send(c context.Context, r *ChunkPb.ChunkRequest) (*ChunkPb.ChunkResponse, error) {

	return &ChunkPb.ChunkResponse{}, nil
}
