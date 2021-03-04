package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultBadRequest(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusBadRequest, Exception{
		"message": fmt.Sprintf("[%s] %s", DefaultBadRequestException, "invalid request"),
	})
}
