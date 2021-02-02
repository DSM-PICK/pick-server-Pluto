package web

import "github.com/gin-gonic/gin"

func Main() {
	router := gin.New()
	router.Use(gin.Recovery())

	activityRouter := router.Group("/activity")
	{
		activityRouter.POST("/", setActivities)
	}

	router.Run()
}
