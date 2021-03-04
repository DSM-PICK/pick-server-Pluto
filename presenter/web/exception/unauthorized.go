package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InvalidAuthorization(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusUnauthorized, Exception{
		"message": fmt.Sprintf("[%s] %s", InvalidAuthorizationException, "invalid authorization"),
	})
}
