/*
 * @Author: zzh weiersili2021@163.com
 * @Date: 2024-10-25 11:12:31
 * @LastEditors: zzh weiersili2021@163.com
 * @LastEditTime: 2024-10-25 15:18:01
 * @FilePath: /Goproject/Gin/GraduationDesign/Router/router.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package router

import (
	common "github.com/GraduationDesign/Common"
	controller "github.com/GraduationDesign/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "pong",
		})
	})
	r.POST("/user/register", controller.Register)
	r.POST("/user/login", controller.Login)
	common.UploadSingleFile("/user/UploadFile", r)
	common.UploadMultiFiles("/user/UploadFiles", r)
	return r
}
