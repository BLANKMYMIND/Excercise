package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	router := gin.Default()
	router.POST("/ti/setcookie", func(context *gin.Context) {
		context.SetCookie("test", "123456", 30, "/", "localhost", http.SameSiteNoneMode, false, false)
		context.JSON(http.StatusOK, gin.H{
			"state": "success",
			"cookie": "123456",
		})
	})
	router.POST("/ti/getcookie", func(context *gin.Context) {
		cookie, err := context.Cookie("test")
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"state": "error",
				"cookie": cookie,
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"state": "success",
			"cookie": cookie,
		})
	})
	router.Run(":8899")
}
