package web

import (
	"excel-import/presenter/parser"
	"excel-import/usecase"
	"excel-import/usecase/exception"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
)

func setActivities(context *gin.Context) {
	fmt.Println()
	formData, e := context.FormFile("file")
	if e != nil || filepath.Ext(strings.ToLower(formData.Filename)) != ".xlsx" {
		exception.DefaultBadRequest(context)
		return
	}

	file, _ := formData.Open()
	defer file.Close()

	usecase.SetActivities(parser.Activity(file, formData.Size))

	context.Status(http.StatusCreated)
}
