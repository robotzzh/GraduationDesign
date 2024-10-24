/*
 * @Author: zzh weiersili2021@163.com
 * @Date: 2024-10-20 19:56:52
 * @LastEditors: zzh weiersili2021@163.com
 * @LastEditTime: 2024-10-24 21:19:51
 * @FilePath: /Goproject/Gin/Project/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	common "github.com/GraduationDesign/Common" // 注意这个导入的下划线，表示仅初始化驱动
	controller "github.com/GraduationDesign/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	//router := gin.Default()
	common.InitDB()
	db := common.GetDB()
	defer db.Close()
	//router.Run(":8080")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "pong",
		})
	})
	r.POST("/user/register", controller.Register)
	r.Run(":8080")
}
