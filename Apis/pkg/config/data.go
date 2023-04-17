package config

import (
	"github.com/soheil-89/OMPFinex/Apis/pkg/config/environment"
)

var Env = newEnvData()

type envData struct {
	App         app
	Db          db
	Grpc        grpc
	PublicFiles publicFiles
}
type publicFiles struct {
	Path string
}
type app struct {
	TimeZone string
	Port     string
	Url      string
	GinMode  string
}

type db struct {
	Sqlite
}
type Sqlite struct {
	Name string
}
type grpc struct {
	Uri  string
	Port string
}

func newEnvData() *envData {
	env := environment.Load()
	return &envData{
		App: app{
			TimeZone: env.Get("app.TimeZone").(string),
			Port:     env.Get("app.Port").(string),
			Url:      env.Get("app.Url").(string),
			GinMode:  env.Get("app.GIN_MODE").(string),
		},
		Db: db{
			Sqlite{
				Name: env.Get("db.Sqlite.Name").(string),
			},
		},
		Grpc: grpc{
			Uri:  env.Get("grpc.Uri").(string),
			Port: env.Get("grpc.Port").(string),
		},
		PublicFiles: publicFiles{
			Path: env.Get("publicFiles.path").(string),
		},
	}
}
