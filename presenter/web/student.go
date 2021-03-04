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

func setStudents(context *gin.Context) {
	formData, e := context.FormFile("file")
	if e != nil || filepath.Ext(strings.ToLower(formData.Filename)) != ".xlsx" {
		exception.DefaultBadRequest(context) ; return
	}

	usecase.SetStudent(excel.ParseStudents(formData))

	context.Status(http.StatusCreated)
}
