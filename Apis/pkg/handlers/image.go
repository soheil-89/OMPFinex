package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/soheil-89/OMPFinex/Apis/pkg/dbConnection"
	"github.com/soheil-89/OMPFinex/Apis/pkg/models"
	json "github.com/soheil-89/OMPFinex/Apis/pkg/services/response"
	"github.com/soheil-89/OMPFinex/Apis/pkg/validation"
	"net/http"
)

func Register(ctx *gin.Context) {
	var validate validation.Register
	err := ctx.BindJSON(&validate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, json.Json("error", err))
		return
	}
	var imageModel models.Image
	models.Make(&imageModel, validate)
	dbConnection.DbSqlite.Create(&imageModel)
	ctx.JSON(http.StatusCreated, imageModel)
}

func Chunk(ctx *gin.Context) {
	var validate validation.Chunk
	err := ctx.BindJSON(&validate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, json.Json("error", err))
		return
	}
	var imageModel models.Image
	models.Make(&imageModel, validate)
	dbConnection.DbSqlite.Create(&imageModel)
	ctx.JSON(http.StatusCreated, imageModel)
}

func Download(ctx *gin.Context) {
	image := models.Image{
		Sha256:    "12",
		Size:      12,
		ChunkSize: 12,
		Data:      "12",
	}
	dbConnection.DbSqlite.Create(&image)
}
