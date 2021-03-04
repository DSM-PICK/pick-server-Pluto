package web

import (
	"github.com/gin-gonic/gin"
)

func Main() {
	router := gin.New()

	router.Use(gin.Recovery())

	plutoRouter := router.Group("/pluto")
	{
		plutoRouter.POST("/activity", setActivities)
		plutoRouter.POST("/student", setStudents)
	}

	router.Run()
}
