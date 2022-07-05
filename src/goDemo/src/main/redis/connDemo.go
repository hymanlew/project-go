package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/*
golang 操作 redis 的客户端包有多个，比如 redigo、go-redis，目前 Star 最多的莫属 redigo。
安装第三方开源 redis 库，在 GOPATH 路径下项目的路径下执行以下命令（以下是两个工具包，任选一个即可）：
go get github.com/garyburd/redigo/redis
go get github.com/go-redis/redis/v8
go get gopkg.in/redis.v4

redigo 文档：https://godoc.org/github.com/garyburd/redigo/redis

Conn 接口是与 Redis 协作的主要接口，可使用 Dial, DialWithTimeout, NewConn 函数来创建连接。最后必须调用 Close 函数关闭。

*/

func demo() {
	//1. 链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()

	//2. 向 redis 写入数据 string [key-val]
	_, err = conn.Do("Set", "name", "tomjerry猫猫")
	if err != nil {
		fmt.Println("set  err=", err)
		return
	}
	//设置 key 过期时间
	_, err = conn.Do("expire", "name", 10) //10秒过期
	if err != nil {
		fmt.Println("set expire error: ", err)
		return
	}

	//3. 向 redis 读取数据 string [key-val]
	//因为 do 函数返回的是 interface{} 接口，所以当查出数据后，需要进行转换。但不能如下直接转换
	//nameString := r.(string)
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set  err=", err)
		return
	}
	fmt.Println("get 数据 = ", r)

	//4. hash 数据操作
	_, err = conn.Do("HSet", "user01", "name", "john")
	if err != nil {
		fmt.Println("hset  err=", err)
		return
	}
	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("hset  err=", err)
		return
	}
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}
	r2, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}
	fmt.Printf("hget r1=%v r2=%v \n", r1, r2)

	//5. hash 批量数据操作
	_, err = conn.Do("HMSet", "user02", "name", "john", "age", 19)
	if err != nil {
		fmt.Println("HMSet  err=", err)
		return
	}
	r, err = redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("hget  err=", err)
		return
	}
	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}
