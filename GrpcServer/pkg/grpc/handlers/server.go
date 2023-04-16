package grpcHandlers

import (
	"fmt"
	"github.com/soheil-89/OMPFinex/GrpcServer/pkg/config"
	"github.com/soheil-89/OMPFinex/GrpcServer/pkg/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func ServerRun() {
	lis, err := net.Listen("tcp", ":"+config.Env.Grpc.Port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
	s := grpc.NewServer()
	ChunkPb.RegisterChunkServer(s, &Server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
