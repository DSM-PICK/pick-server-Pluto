package web

import (
	"github.com/gin-gonic/gin"
	"os"
	"pluto/presenter/web/middleware"
)

func Main() {
	var e error
	middleware.Logger.Out, e = os.OpenFile(os.Getenv("LOG_PATH"),
										   os.O_CREATE|os.O_WRONLY|os.O_APPEND,
										   0666)
	if e != nil { panic(e) }

	router := gin.New()

	router.Use(middleware.Log())
	router.Use(middleware.Recover())

	plutoRouter := router.Group("/pluto")
	plutoRouter.Use(middleware.Auth())
	{
		plutoRouter.POST("/activity", setActivities)
		plutoRouter.POST("/student", setStudents)
	}

	router.Run()
}
