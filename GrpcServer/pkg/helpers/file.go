package helpers

import (
	"github.com/soheil-89/OMPFinex/GrpcServer/pkg/config"
	"os"
)

func CreateFile(file string, fileName string) error {
	f, err := os.Create(config.Env.PublicFiles.Path + fileName + ".txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	_, err = f.WriteString(file)
	if err != nil {
		return err
	}
	return nil
}
