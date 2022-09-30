package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHealth(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Service Up")
}
