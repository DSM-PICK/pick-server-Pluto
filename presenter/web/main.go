package web

import (
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx/v3"
	"log"
	"os"
	"path/filepath"
	webException "pluto/presenter/web/exception"
	"pluto/presenter/web/middleware"
	parserException "pluto/usecase/parser"
	"strings"
)

func Main() {
	var e error
	middleware.Logger.Out, e = os.OpenFile(os.Getenv("LOG_PATH"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if e != nil {
		panic(e)
	}

	router := gin.New()

	router.Use(middleware.Log())
	router.Use(middleware.Recover())
	router.Use(middleware.Cors())

	plutoRouter := router.Group("/pluto")
	{
		plutoRouter.POST("/activity", setActivities).Use(middleware.Auth("SECRET_KEY"))
		plutoRouter.POST("/student", setStudents).Use(middleware.Auth("ADMIN_SECRET_KEY"))
	}

	log.Fatal(router.Run())
}

func getXlsxFromFormData(context *gin.Context) (*xlsx.File, error) {
	formData, e := context.FormFile("file")
	if e != nil {
		webException.DefaultBadRequest(context)
		return nil, parserException.InvalidFile
	}

	if filepath.Ext(strings.ToLower(formData.Filename)) != ".xlsx" {
		return nil, parserException.InvalidFile
	}

	file, _ := formData.Open()
	defer file.Close()

	workbook, _ := xlsx.OpenReaderAt(file, formData.Size)
	return workbook, nil
}
