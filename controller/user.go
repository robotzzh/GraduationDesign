/*
 * @Author: zzh weiersili2021@163.com
 * @Date: 2024-10-24 21:06:13
 * @LastEditors: zzh weiersili2021@163.com
 * @LastEditTime: 2024-10-25 11:30:57
 * @FilePath: /Goproject/Gin/GraduationDesign/controller/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"net/http"

	common "github.com/GraduationDesign/Common"
	model "github.com/GraduationDesign/Model"
	response "github.com/GraduationDesign/Response"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // 注意这个导入的下划线，表示仅初始化驱动
)

/*
测试部分
curl -X POST http://localhost:8080/user/register \
-H "Content-Type: application/json" \
-d '{
  "Username": "testuser",
  "Mobile": "1234567890",
  "Password": "testpassword"
}'
*/

func Register(c *gin.Context) {
	db := common.GetDB()

	reqUser := model.User{}
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	username := reqUser.Username
	mobile := reqUser.Mobile
	password := reqUser.Password

	newUser := &model.User{
		Username: username,
		Mobile:   mobile,
		Password: password,
	}

	Result := db.Create(newUser)
	if Result != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 114514,
			"msg":  Result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

/*
测试部分
正确测试
curl -X POST http://localhost:8080/user/login \
-H "Content-Type: application/json" \
-d '{
  "Username": "testuser",
  "Mobile": "1234567890",
  "Password": "testpassword"
}'
错误测试
curl -X POST http://localhost:8080/user/login \
-H "Content-Type: application/json" \
-d '{
  "Username": "test",
  "Mobile": "1234567890",
  "Password": "testpassword"
}'

curl -X POST http://localhost:8080/user/login \
-H "Content-Type: application/json" \
-d '{
  "Username": "testuser",
  "Mobile": "1234567890",
  "Password": "testpasd"
}'
*/

func Login(c *gin.Context) {
	db := common.GetDB()

	reqUser := model.User{}
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		response.Response(c, http.StatusBadRequest, 114514, "未发现用户", gin.H{})
		return
	}

	username := reqUser.Username
	mobile := reqUser.Mobile
	password := reqUser.Password

	newUser := &model.User{
		Username: username,
		Mobile:   mobile,
		Password: password,
	}

	var findUser model.User
	if err := db.Where("username = ?", newUser.Username).First(&findUser).Error; err != nil {
		response.Response(c, http.StatusOK, 114514, "用户不存在", gin.H{})
	} else {
		if findUser.Password == newUser.Password {
			response.Success(c, "登录成功", gin.H{})
		} else {
			response.Response(c, http.StatusOK, 114514, "密码错误", gin.H{})
		}
	}
}
