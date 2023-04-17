package main

import (
	"github.com/soheil-89/OMPFinex/GrpcServer/pkg/dbConnection"
	grpcHandlers "github.com/soheil-89/OMPFinex/GrpcServer/pkg/grpc/handlers"
	"github.com/soheil-89/OMPFinex/GrpcServer/pkg/models"
)

func main() {

	dbConnection.ConnectSqlite()
	_ = dbConnection.DbSqlite.AutoMigrate(&models.Image{})
	grpcHandlers.ServerRun()

}
