package dbConnection

import (
	"fmt"
	"github.com/soheil-89/OMPFinex/GrpcServer/pkg/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var DbSqlite *gorm.DB

func ConnectSqlite() {
	db, err := gorm.Open(sqlite.Open(config.Env.Db.Sqlite.Name+".db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	DbSqlite = db
}
