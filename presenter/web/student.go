package web

import (
	"excel-import/presenter/parser/excel"
	"excel-import/presenter/web/exception"
	"excel-import/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
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
