package main

import (
	"database/sql"
	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func getHost() (host string) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Panic(err)
	}
	defer c.Close()
	res, err := c.Do("get", "sqlHost")
	if err != nil {
		log.Panic(err)
	}
	log.WithFields(log.Fields{
		"word": res,
	}).Info()
	host = string(res.([]uint8))
	return
}

func getPassword() (password string) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Panic(err)
	}
	res, err := c.Do("get", "sqlPassword")
	if err != nil {
		log.Panic(err)
	}
	log.WithFields(log.Fields{
		"word": res,
	}).Info()
	// 注意！ redis 返回的是字节切片 []uint8，需要先断言再转 string 后才能赋值给返回值
	password = string(res.([]uint8))
	return
}

func main() {
	// 先从本地 redis 里拿到敏感信息
	host := getHost()
	password := getPassword()

	// pq 连接使用键值对的字符串形式传参数
	// https://godoc.org/github.com/lib/pq
	connStr := "host=" + host + " dbname=application user=postgres password=" + password + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}
	log.WithFields(log.Fields{
		"action": "database",
		"state":  "connected",
	}).Info()

	// pq 预处理不使用 ? 作保留符，而是使用 $1 ~ $n 作保留符
	stmt, err := db.Prepare("SELECT uid from account WHERE phone=$1")
	if err != nil {
		log.Panic(err)
	}
	res, err := stmt.Query("15570877777")
	if err != nil {
		log.Panic(err)
	}
	// 只需取头个的情况
	if res.Next() {
		var uid int64
		res.Scan(&uid)
		log.WithFields(log.Fields{
			"action": "queryuid",
			"state":  "found",
			"uid":    uid,
		}).Info()
	} else {
		log.WithFields(log.Fields{
			"action": "queryid",
			"state":  "miss",
		}).Warn()
	}
	defer db.Close()
}
