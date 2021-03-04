package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func Log() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()

		context.Next()

		entry := logrus.WithFields(logForm(context, time.Now().Sub(start)))

		if status := context.Writer.Status() ; http.StatusBadRequest > status && status >= http.StatusOK {
			entry.Info()
		} else if http.StatusInternalServerError > status && status >= 400 {
			entry.Warn()
		}
	}
}

func logForm(context *gin.Context, duration time.Duration) logrus.Fields {
	bodyBytes, _ := ioutil.ReadAll(context.Request.Body)
	return logrus.Fields{
		"client_ip": context.ClientIP(),
		"duration": duration,
		"method": context.Request.Method,
		"path": context.Request.RequestURI,
		"status": context.Writer.Status(),
		"body": string(bodyBytes),
		"formData": context.Request.MultipartForm,
	}
}
