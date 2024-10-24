/*
 * @Author: zzh weiersili2021@163.com
 * @Date: 2024-10-24 16:48:06
 * @LastEditors: zzh weiersili2021@163.com
 * @LastEditTime: 2024-10-24 17:13:20
 * @FilePath: /Goproject/Gin/GraduationDesign/Common/DB.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package common

import (
	"fmt"

	model "github.com/GraduationDesign/Model"
	_ "github.com/go-sql-driver/mysql" // 注意这个导入的下划线，表示仅初始化驱动
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	diverName := "mysql"
	host := "localhost"
	port := "3306"
	username := "root"
	password := "12345678"
	database := "GD"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(diverName, args)
	db.AutoMigrate(&model.User{})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
