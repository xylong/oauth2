package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrorHandler 错误处理
func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err:=recover();err!=nil {
				context.JSON(http.StatusBadRequest,gin.H{
					"msg":err,
				})
			}
		}()
		context.Next()
	}
}
