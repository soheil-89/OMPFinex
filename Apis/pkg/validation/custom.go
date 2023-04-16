package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/soheil-89/OMPFinex/Apis/pkg/dbConnection"
	"strings"
)

func Run() {
	v, _ := binding.Validator.Engine().(*validator.Validate)
	_ = v.RegisterValidation("exist_in_db", ExistInDb)

}

func ExistInDb(fl validator.FieldLevel) bool {
	key := strings.Split(fl.Param(), ".")
	value := fl.Field().String()
	var count int
	dbConnection.DbSqlite.Raw("SELECT COUNT(id) FROM "+key[0]+" WHERE "+key[1]+" = ?", value).Scan(&count)
	if count > 0 {
		return false
	}
	return true
}
