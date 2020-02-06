package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func exampleMySQL() {
	// MySQL
	// 连接数据库，得到 *DB 类型的 db 对象
	db, err := sql.Open("mysql", "root:password@/test?charset=utf8")
	checkErr(err, 1)

	// 预编译 SQL，得到 *stmt 类型的 stmt 对象，此时会获得一个 单线程使用 的连接
	stmt, err := db.Prepare("INSERT INTO user set username=?, sex=?")
	checkErr(err, 2)

	// 预编译 SQL 执行，得到 Result 类型的 res 对象
	res, err := stmt.Exec("bye-bye", 2)
	checkErr(err, 3)

	// 插入的 id
	fmt.Println(res.LastInsertId())
	// 影响的行数
	fmt.Println(res.RowsAffected())

	stmt, err = db.Prepare("UPDATE user set sex=? where uid=?")
	checkErr(err, 4)

	res, err = stmt.Exec(10, 2)
	checkErr(err, 5)

	// 插入的 id （此时只是更新，因此值为 0）
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())

	rows, err := db.Query("SELECT * FROM user where 1;")
	checkErr(err, 6)

	// 提示了各字段顺序
	fmt.Println(rows.Columns())

	for rows.Next() {
		var sex int
		var username string
		var uid int
		// 注意！ Scan 是按照字段顺序依次赋值，与变量名无关（例子如下）
		err = rows.Scan(&sex, &username, &uid)
		checkErr(err, 7)
		fmt.Println(sex, username, uid)
	}
}

func exampleRedis() {
	// Redis
	// 新建连接池
	redisHost := ":6379"
	var Pool *redis.Pool
	Pool = &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	// 连接关闭逻辑
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()

	// 实操环节
	conn := Pool.Get()
	defer conn.Close()
	var data string
	data, err := redis.String(conn.Do("GET", "test"))
	fmt.Println(data, err)

	data, err = redis.String(conn.Do("SET", "do", "now"))
	fmt.Println(data, err)
}

func main() {
	// exampleMySQL()
	exampleRedis()
}

func checkErr(err error, place int) {
	if err != nil {
		panic(place)
		panic(err)
	}
}
