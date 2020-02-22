package main

import (
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/big"
	"net/http"
	"path"
	"time"
)

func main() {
	router := gin.Default()
	// 使用路径参数
	// : 标识 该路由的路径不包括 /user/ or /user
	// * 标识 该路由的路径将包括 /user/ or /user (当 /user 在其他路由中不匹配时才会跳到这里)
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s", name)
	})
	// 路径冲突会有提示，下面这个路由因冲突无法添加（报 panic）
	//router.GET("/user/john", func(c *gin.Context) {
	//	c.String(http.StatusOK, "Sorry, it's Tom.", )
	//})

	// 使用字符串参数 (传统 GET)
	router.GET("/welcome", func(c *gin.Context) {
		// defaultQ 取出参数 若无，返回默认设置的值
		first := c.DefaultQuery("first", "anonymous")
		// Q 取出参数 若无，返回“”
		last := c.Query("last")
		// 使用 c.Request 可较直接详细查取 request 中的内容
		c.String(http.StatusOK, "Hello, "+first+" "+last+".")
	})

	// 使用 form-data 参数
	router.POST("form_post", func(c *gin.Context) {
		// PF 取参
		message := c.PostForm("message")
		// DefaultPF 取参并设默认值
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// 字符串参数 和 form-data 参数的 map 形式
	// 请求格式： /post?names[aa]=bb&names[cc]=dd 和 form-data 的相似形式
	router.POST("/post", func(c *gin.Context) {
		ids := c.PostFormMap("ids")
		names := c.QueryMap("names")
		c.String(http.StatusOK, "ids: %v, names: %v", ids, names)
		// 输出 ids: map[], names: map[aa:bb cc:dd]
	})

	// 上传单个文件
	router.POST("/upload", func(c *gin.Context) {
		// 获取解析文件
		file, err := c.FormFile("file")
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"action": "file upload fail",
				"error":  err,
			}).Warn()
			c.String(http.StatusBadRequest, "Sorry, file-upload process is FAIL.")
			// 错误处理完成后记得要 return
			return
		}
		logrus.WithFields(logrus.Fields{
			"fileName": file.Filename,
		}).Info()
		t := time.Now()
		s := fmt.Sprintf("%v_%v_%v_%v_%v_%v", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
		// 获取文件后缀 (Ext 方法获取的后缀带 . , 不用自己加)
		suffix := path.Ext(file.Filename)
		// 设定目标地址及文件名
		dst := "./test/" + s + suffix
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"action": "file upload fail",
				"error":  err,
			}).Warn()
			c.String(http.StatusBadRequest, "Sorry, file-upload process is FAIL.")
			return
		}
		c.String(http.StatusOK, "%s has been uploaded!", file.Filename)
	})

	// 上传多个文件
	router.POST("/upload_multiple", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"action": "file upload fail",
				"error":  err,
			}).Warn()
			c.String(http.StatusBadRequest, "Sorry, file-upload process is FAIL.")
			// 错误处理完成后记得要 return
			return
		}
		// 这些文件的 key 都要是 files (form-data 的 key 是可重复的)
		files := form.File["files"]
		max := new(big.Int).SetUint64(uint64(1024))
		// 循环 files 来一个个存 file
		for _, file := range files {
			r, _ := rand.Int(rand.Reader, max)
			t := time.Now()
			// 文件名加上一个随机数防止重复
			s := fmt.Sprintf(
				"%v_%v_%v_%v_%v_%v_%v",
				t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second(),
				r,
			)
			// 获取文件后缀 (Ext 方法获取的后缀带 . , 不用自己加)
			suffix := path.Ext(file.Filename)
			// 设定目标地址及文件名
			dst := "./test/" + s + suffix
			if err := c.SaveUploadedFile(file, dst); err != nil {
				logrus.WithFields(logrus.Fields{
					"action": "file upload fail",
					"error":  err,
				}).Warn()
				c.String(http.StatusBadRequest, "Sorry, file-upload process is FAIL.")
				return
			}
		}
		c.String(http.StatusOK, "success")
	})

	// 路由分组
	// 目的是加强逻辑间关系，提升路由匹配速度
	//
	// Simple group: v1
	//v1 := router.Group("/v1")
	//{
	//	v1.POST("/login", loginEndpoint)
	//	v1.POST("/submit", submitEndpoint)
	//	v1.POST("/read", readEndpoint)
	//}
	//
	//v2 := router.Group("/v2")
	//{
	//	v2.POST("/login", loginEndpoint)
	//	v2.POST("/submit", submitEndpoint)
	//	v2.POST("/read", readEndpoint)
	//}
	//
	// 此时 需要访问 /v1/login /v2/login 才能匹配到相应 handler

	// 使用中间件进行验证操作
	auth := router.Group("/auth", func(c *gin.Context) {
		token := c.PostForm("token")
		if token != "123" {
			//c.JSON(http.StatusBadRequest, gin.H{
			//	"state": "auth_fail",
			//})
			//// 验证错误时，直接 abort 返回
			//c.Abort()
			// 也可以直接使用
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"state": "auth_fail",
			})
		} else {
			c.Set("type", "admin")
		}
	})
	// 这只是一个代码块，不是任何函数的主体
	{
		auth.POST("/t", func(c *gin.Context) {
			state, _ := c.Get("type")
			logrus.WithFields(logrus.Fields{
				"state": state,
			}).Info()
			// 上面使用的 set get 不影响原请求的信息
			userState := c.PostForm("type")
			//userState := c.GetHeader("type")
			logrus.WithFields(logrus.Fields{
				"user_state": userState,
			}).Info()
			c.JSON(http.StatusOK, gin.H{
				"hey":   "true",
				"state": state,
			})
		})
	}
	// 静态服务 与 http 的 fileServer 不同，单 /test 页面不显示
	router.Static("/test", "./test")

	// 运行时加载路由是允许的
	//router.GET("/active", func(c *gin.Context) {
	//	c.String(http.StatusOK, "i am dead.")
	//})
	// 但至今没找到运行时删路由的方法...
	go func() {
		time.Sleep(6 * time.Second)
		router.GET("/active", func(c *gin.Context) {
			c.String(http.StatusOK, "i am alive.")
		})
	}()

	router.Run(":8000")
}
