package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"pluto/presenter/parser/excel"
	"pluto/presenter/web/exception"
	"pluto/usecase"
	"strings"
)

func setActivities(context *gin.Context) {
	formData, e := context.FormFile("file")
	if e != nil || filepath.Ext(strings.ToLower(formData.Filename)) != ".xlsx" {
		exception.DefaultBadRequest(context) ; return
	}

	usecase.SetActivities(excel.ParseActivities(formData))

	context.Status(http.StatusCreated)
}
