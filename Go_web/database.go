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

func once(Pool *redis.Pool, i int, ch chan bool) {
	conn := Pool.Get()
	_, err := conn.Do("lpush", "goods&10000", i)
	if err != nil {
		fmt.Println(err)
	}
	conn.Close()
	ch <- true
}

func concurrencyRedis(c2 redis.Conn, Pool *redis.Pool) {
	// 并发测试 插入10000个数据
	tOri := time.Now()
	l := make([]int, 10000)
	ch := make(chan bool, 10000)
	for i := range l {
		go once(Pool, i, ch)
	}
	k := 0
	for range ch {
		if k++; k >= 10000 {
			break
		}
	}
	d, _ := redis.Int64s(c2.Do("lrange", "goods&10000", "0", "10000"))
	tClo := time.Now()
	t := tClo.Sub(tOri)
	fmt.Println(len(d))
	fmt.Println(t)
	c2.Do("del", "goods&10000")
}

func exampleRedis() {
	// Redis
	// 新建连接池
	redisHost := ":6379"
	var Pool *redis.Pool
	Pool = &redis.Pool{
		// 高并发必备的两个参数，最优设置（非最大设置）可使速度飞起来
		// 优化前 1.7s ~ 6.8s 浮动，优化后平均 0.42s
		// MaxActive 决定最大链接数，保证不超过win的最大链接数
		// Wait true 时， 开启等待模式： 无可用连接时等待
		MaxActive: 400,
		Wait:      true,

		MaxIdle:     30,
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

	c2 := Pool.Get()
	defer c2.Close()
	d, err := redis.Int64s(c2.Do("lrange", "goods&10086", "0", "10"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)

	// 并发测试
	// concurrencyRedis(c2, Pool)

	// 设置过期时间
	cc := Pool.Get()
	defer cc.Close()
	res, err := redis.String(cc.Do("set", "t", time.Now()))
	fmt.Println(res, err)

	res, err = redis.String(cc.Do("get", "t"))
	fmt.Println(res, err)

	i, err := redis.Int64(cc.Do("expire", "t", 3))
	fmt.Println(i, err)

	time.Sleep(time.Second * 1)

	i, err = redis.Int64(cc.Do("ttl", "t"))
	fmt.Println(i, err)

	i, err = redis.Int64(cc.Do("pttl", "t"))
	fmt.Println(i, err)

	// 注意 重新设定值会清除先前设定的时间 需要重新设定
	// res, err = redis.String(cc.Do("set", "t", "sss"))
	// fmt.Println(res, err)

	time.Sleep(time.Second * 2)

	g, err := redis.String(cc.Do("get", "t"))
	fmt.Println(g, err)
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
