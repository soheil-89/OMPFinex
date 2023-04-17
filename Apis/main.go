package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/soheil-89/OMPFinex/Apis/pkg/config"
	"github.com/soheil-89/OMPFinex/Apis/pkg/handlers"
	"os"
)

func main() {
	r := gin.Default()
	r.POST("/image", handlers.Register)
	r.POST("/image/:sha256/chunks", handlers.Chunk)
	r.GET("/image/:sha256", handlers.Download)
	err := r.Run(":" + config.Env.App.Port)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
