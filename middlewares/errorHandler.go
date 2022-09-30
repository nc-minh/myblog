package middlewares

import (
	"fmt"
	"myblog/utils/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	for _, err := range ctx.Errors {
		logger.Debug(fmt.Sprintf("Error: %s", err.Error()))
	}

	ctx.JSON(http.StatusInternalServerError, "Error")
}
