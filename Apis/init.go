package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soheil-89/OMPFinex/Apis/pkg/config"
	"github.com/soheil-89/OMPFinex/Apis/pkg/dbConnection"
	"github.com/soheil-89/OMPFinex/Apis/pkg/models"
	"github.com/soheil-89/OMPFinex/Apis/pkg/validation"
	"time"
)

func init() {
	//db connection
	dbConnection.ConnectSqlite()
	_ = dbConnection.DbSqlite.AutoMigrate(&models.Image{})

	//custom validation
	validation.Run()

	//Gin Set Mode
	gin.SetMode(config.Env.App.GinMode)

	//Time Zone
	loc, _ := time.LoadLocation(config.Env.App.TimeZone)
	time.Local = loc
}
