package api

import (
	"os"

	"github.com/gin-gonic/gin"
)

func GetHomeWork(c *gin.Context) {
	version := os.Getenv("VERSION")

	requestWriteMap := make(map[string]interface{})
	clientIP, _ := c.RemoteIP()

	for k := range c.Request.Header {
		requestWriteMap[k] = c.Request.Header[k]
	}

	requestWriteMap["versionEnv"] = version
	requestWriteMap["clientIP"] = clientIP

	JSONResponse(c, requestWriteMap)
}
