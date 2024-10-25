/*
 * @Author: zzh weiersili2021@163.com
 * @Date: 2024-10-20 19:56:52
 * @LastEditors: zzh weiersili2021@163.com
 * @LastEditTime: 2024-10-25 11:15:09
 * @FilePath: /Goproject/Gin/Project/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	common "github.com/GraduationDesign/Common"
	router "github.com/GraduationDesign/Router"
	"github.com/gin-gonic/gin"
)

func main() {

	//router := gin.Default()
	common.InitDB()
	db := common.GetDB()
	defer db.Close()
	//router.Run(":8080")
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")
}
