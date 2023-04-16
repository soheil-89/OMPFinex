package grpcPkg

import (
	"context"
	"github.com/soheil-89/OMPFinex/Apis/pkg/config"
	ChunkPb "github.com/soheil-89/OMPFinex/Apis/pkg/grpc/proto"
	"github.com/soheil-89/OMPFinex/Apis/pkg/validation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Data struct {
	*validation.Chunk
}

func (d Data) Request(sha256 string) (bool, string) {
	conn, err := grpc.Dial(config.Env.Grpc.Uri+":"+config.Env.Grpc.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, err.Error()
	}
	defer conn.Close()
	c := ChunkPb.NewChunkClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Send(ctx, &ChunkPb.ChunkRequest{
		Sha256: sha256,
		Id:     d.Id,
		Size:   d.Size,
		Data:   d.Data,
	})
	return r.Receive, err.Error()
}
