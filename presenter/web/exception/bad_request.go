package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultBadRequest(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusBadRequest, map[string]string {
		"message": fmt.Sprintf("[%s] %s", defaultBadRequest, "invalid request"),
	})
}
