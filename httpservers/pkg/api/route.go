package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	// Creates a gin restful with default middleware:
	router := gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	core := router.Group("/api")
	{
		core.GET("/homework", GetHomeWork)
	}

	router.Handle("GET", "/localhost/healthz", Health)

	return router
}
