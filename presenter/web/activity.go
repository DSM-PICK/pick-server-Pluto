package web

import (
	"pluto/presenter/parser/excel"
	"pluto/presenter/web/exception"
	"pluto/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
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
