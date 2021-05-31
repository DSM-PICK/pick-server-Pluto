package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pluto/presenter/web/exception"
	"pluto/usecase"
	"pluto/usecase/parser/excel"
)

func setActivities(context *gin.Context) {
	file, e := getXlsxFromFormData(context)
	if e != nil {
		return
	}

	e = usecase.SetActivities(excel.ParseActivities(file))
	if e != nil {
		exception.DefaultBadRequest(context)
		return
	}

	context.Status(http.StatusCreated)
}
