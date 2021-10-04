package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is a uniform output format for the API
type Response struct {
	// Code 0 means success, non-0 means failure
	Code int64 `json:"code"`

	// The function of this field overlaps with Code
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONResponse(c *gin.Context, data interface{}) {
	resp := Response{
		Code:    0,
		Success: true,
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}

func JSONResponseError(c *gin.Context, msg string) {
	resp := Response{
		Code:    1,
		Success: false,
		Message: msg,
	}
	c.JSON(http.StatusOK, resp)
}

func Health(c *gin.Context) {
	JSONResponse(c, nil)
}
