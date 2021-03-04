package web

import (
	"github.com/gin-gonic/gin"
	"os"
	"pluto/presenter/web/middleware"
)

func Main() {
	middleware.Logger.Out, _ = os.OpenFile("./resources/log/pluto.log",
										   os.O_CREATE|os.O_WRONLY|os.O_APPEND,
										   0666)

	router := gin.New()

	router.Use(middleware.Log())
	router.Use(middleware.Recover())

	plutoRouter := router.Group("/pluto")
	{
		plutoRouter.POST("/activity", setActivities)
		plutoRouter.POST("/student", setStudents)
	}

	router.Run()
}
