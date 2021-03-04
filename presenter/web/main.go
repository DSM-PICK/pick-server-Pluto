package web

import (
	"github.com/gin-gonic/gin"
	"pluto/presenter/web/middleware"
)

func Main() {
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
