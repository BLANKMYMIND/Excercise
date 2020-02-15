package main

import "github.com/gin-gonic/gin"
import log "github.com/sirupsen/logrus"

func main() {
	r := gin.Default()
	// 路径前有 /
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()
	// gin 主体服务 正常运行情况下不会执行到 err 处理的内容
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Panic()
	} else {
		log.WithFields(log.Fields{
			"status": "success",
		}).Info()
	}
}
