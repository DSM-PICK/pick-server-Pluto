package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Recover() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()

		defer func() {
			if r := recover(); r != nil {
				context.Status(http.StatusInternalServerError)
				logrus.WithFields(logForm(context, time.Now().Sub(start))).Error(r)
			}
		}()

		context.Next()
	}
}
