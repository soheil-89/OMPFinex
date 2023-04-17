package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/soheil-89/OMPFinex/Apis/pkg/config"
	"github.com/soheil-89/OMPFinex/Apis/pkg/dbConnection"
	"github.com/soheil-89/OMPFinex/Apis/pkg/endpoint"
	grpcPkg "github.com/soheil-89/OMPFinex/Apis/pkg/grpc/pkg"
	"github.com/soheil-89/OMPFinex/Apis/pkg/helpers"
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
	helpers.Inject(&imageModel, validate)
	dbConnection.DbSqlite.Create(&imageModel)
	var register endpoint.Register
	helpers.Inject(&register, imageModel)
	ctx.JSON(http.StatusCreated, register)
}

func Chunk(ctx *gin.Context) {
	sha256 := ctx.Param("sha256")
	var validate validation.Chunk
	err := ctx.BindJSON(&validate)
	var image models.Image
	dbConnection.DbSqlite.Where("sha256 = ?", sha256).First(&image)
	if err != nil || image.Sha256 == "" {
		ctx.JSON(http.StatusBadRequest, json.Json("error", err))
		return
	}
	response, err := grpcPkg.NewData(validate).Request(sha256, image.Size)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "we have problem",
		})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func Download(ctx *gin.Context) {
	ctx.File(config.Env.PublicFiles.Path + ctx.Param("sha256") + ".txt")
}
