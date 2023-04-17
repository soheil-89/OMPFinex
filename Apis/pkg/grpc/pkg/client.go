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

type data struct {
	*validation.Chunk
}

func NewData(chunk validation.Chunk) *data {
	return &data{&validation.Chunk{
		Id:   chunk.Id,
		Size: chunk.Size,
		Data: chunk.Data,
	}}
}

func (d *data) Request(sha256 string, size int32) (bool, error) {
	conn, err := grpc.Dial(config.Env.Grpc.Uri+":"+config.Env.Grpc.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, err
	}
	defer conn.Close()
	c := ChunkPb.NewChunkClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Send(ctx, &ChunkPb.ChunkRequest{
		Sha256:    sha256,
		Id:        d.Id,
		Size:      size,
		Data:      d.Data,
		ChunkSize: d.Size,
	})
	return r.Receive, err
}
