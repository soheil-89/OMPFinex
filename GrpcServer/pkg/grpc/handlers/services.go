package grpcHandlers

import (
	"bytes"
	"context"
	"github.com/soheil-89/OMPFinex/GrpcServer/pkg/dbConnection"
	"github.com/soheil-89/OMPFinex/GrpcServer/pkg/grpc/proto"
	"github.com/soheil-89/OMPFinex/GrpcServer/pkg/helpers"
	"github.com/soheil-89/OMPFinex/GrpcServer/pkg/models"
)

type Server struct {
	ChunkPb.UnimplementedChunkServer
}

func (s *Server) Send(c context.Context, r *ChunkPb.ChunkRequest) (*ChunkPb.ChunkResponse, error) {
	var sum int32
	dbConnection.DbSqlite.Raw("SELECT SUM(chunk_size) FROM images WHERE sha256=?", r.Sha256).First(&sum)
	if sum < r.Size {
		image := models.Image{
			Sha256:    r.Sha256,
			ChunkSize: r.ChunkSize,
			Data:      r.Data,
			Number:    r.Id,
		}
		dbConnection.DbSqlite.Create(&image)
	}
	if r.Size <= sum+r.ChunkSize {
		var images []models.Image
		dbConnection.DbSqlite.Order("number asc").Where("sha256 = ?", r.Sha256).Find(&images)
		var allData bytes.Buffer
		for _, image := range images {
			allData.WriteString(image.Data)
		}
		err := helpers.CreateFile(allData.String(), r.Sha256)
		if err != nil {
			return &ChunkPb.ChunkResponse{
				Receive: false,
			}, err
		}
		dbConnection.DbSqlite.Delete(&models.Image{}, "sha256 = ?", r.Sha256)
	}

	return &ChunkPb.ChunkResponse{
		Receive: true,
	}, nil
}
