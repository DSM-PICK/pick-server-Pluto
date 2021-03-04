package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Recover() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()

		defer func() {
			if r := recover(); r != nil {
				context.Status(http.StatusInternalServerError)
				Logger.WithFields(logForm(context, time.Now().Sub(start))).Error(r)
			}
		}()

		context.Next()
	}
}
