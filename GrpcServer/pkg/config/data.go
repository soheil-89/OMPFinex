package config

import (
	"github.com/soheil-89/OMPFinex/GrpcServer/pkg/config/environment"
)

var Env = newEnvData()

type envData struct {
	Grpc        grpc
	Db          db
	PublicFiles publicFiles
}
type publicFiles struct {
	Path string
}
type grpc struct {
	Port string
}
type db struct {
	Sqlite
}
type Sqlite struct {
	Name string
}

func newEnvData() *envData {
	env := environment.Load()
	return &envData{
		Grpc: grpc{
			Port: env.Get("grpc.port").(string),
		},
		Db: db{
			Sqlite{
				Name: env.Get("db.Sqlite.Name").(string),
			},
		},
		PublicFiles: publicFiles{
			Path: env.Get("publicFiles.path").(string),
		},
	}
}
