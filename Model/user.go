/*
 * @Author: zzh weiersili2021@163.com
 * @Date: 2024-10-24 16:59:29
 * @LastEditors: zzh weiersili2021@163.com
 * @LastEditTime: 2024-10-24 17:14:07
 * @FilePath: /Goproject/Gin/GraduationDesign/Model/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

// User
type User struct {
	ID       int    `gorm:"primaryKey"`               //设置为主键
	Username string `gorm:"not null;unique;size:255"` //唯一，不为空
	Password string `gorm:"not null"`
	Mobile   string `gorm:"unique;not null;"`
}
