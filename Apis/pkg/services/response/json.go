package json

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/soheil-89/OMPFinex/Apis/pkg/validation"
)

func Json(key string, data interface{}) gin.H {
	if key == "error" {
		if _, fond := data.(validator.ValidationErrors); fond {
			return gin.H{key: validation.JsonErrorResponse(data.(validator.ValidationErrors))}
		} else {
			return gin.H{key: data}
		}

	}
	return gin.H{key: data}
}

func Metadata(key string, data interface{}, metaData map[string]interface{}) gin.H {
	return gin.H{key: data, "meta_data": metaData}
}
