/*
 * @Author: zzh weiersili2021@163.com
 * @Date: 2024-10-21 23:40:52
 * @LastEditors: zzh weiersili2021@163.com
 * @LastEditTime: 2024-10-24 16:54:24
 * @FilePath: /Goproject/Gin/Project/UploadFile.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package common

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 测试代码
/*
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/zzh/Goproject/testfile.txt" \
  -H "Content-Type: multipart/form-data"
*/
func UploadSingleFile(route string, r *gin.Engine) {
	r.MaxMultipartMemory = 8 << 20

	r.POST(route, func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "File upload error")
			log.Println(http.StatusBadRequest, "File upload error")
			return
		}
		log.Println(file.Filename, "Uploading file")
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		//TODO: finish the logic of uploading process
	})
}

// 测试代码
/*
		curl -X POST http://localhost:8080/upload \
	  -F "upload[]=@/Users/zzh/Goproject/testfile.txt" \
	  -F "upload[]=@/Users/zzh/Goproject/testfile1.txt" \
	  -H "Content-Type: multipart/form-data"
*/

/// @brief Multipart Upload
/// @param route path , gin.Engine

func UploadMultiFiles(route string, router *gin.Engine) {
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件至指定目录
			dst := "./" + file.Filename
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
}
