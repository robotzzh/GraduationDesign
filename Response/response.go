package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 封装返回的代码
func Response(c *gin.Context, httpStatus, code int, msg string, data gin.H) {
	c.JSON(httpStatus, gin.H{"code": code, "msg": msg, "data": data})
}

// Success 成功时返回的代码
func Success(c *gin.Context, msg string, data gin.H) {
	Response(c, http.StatusOK, 200, msg, data)
}
