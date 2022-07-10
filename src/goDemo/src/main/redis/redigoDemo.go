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

//基础连接及使用
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

//全局的连接池 pool
var pool *redis.Pool

//创建 init 函数初始化连接池，当启动程序时就会执行
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,    //最大空闲链接数，即使没有 redis 连接时依然可以保持 N 个空闲的连接，而不被清除，随时处于待命状态
		MaxActive:   16,   //最大链接数，0 表示没有限制
		IdleTimeout: 100,  //最大空闲时间，超过此时间后，空闲连接将被关闭
		Wait:        true, //值为 true 且 MaxActive 参数有限制时，使用 Get 方法获取连接时，会等待一个连接返回给连接池
		Dial: func() (redis.Conn, error) { //初始化链接的代码， 链接哪个ip的redis
			return redis.Dial("tcp", "localhost:6379"), nil
		},
	}
}

func poolDemo() {
	//取出连接，使用完毕后放回连接池
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "汤姆猫~~")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("r=", r)
}

//Pipelining 管道
func pipeDemo() {

	//管道操作可以理解为并发操作，并通过 Send()，Flush()，Receive() 三个方法实现。
	//1. 客户端可以使用 send() 方法一次性向服务器发送一个或多个命令，
	//2. 命令发送完毕时，使用 flush() 方法将缓冲区的命令输入一次性发送到服务器，
	//3. 客户端再使用 Receive() 方法依次按照先进先出的顺序读取所有命令操作结果。

	conn, err := redis.Dial("tcp", "10.1.210.69:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()

	//Send：发送命令至缓冲区
	//Flush：清空缓冲区，将命令一次性发送至服务器
	conn.Send("HSET", "student", "name", "wd", "age", "22")
	conn.Send("HSET", "student", "Score", "100")
	conn.Send("HGET", "student", "age")
	conn.Flush()

	//Recevie：依次读取服务器响应结果，当读取的命令未响应时，该操作会阻塞。
	res1, err := conn.Receive()
	fmt.Printf("Receive res1:%v \n", res1)
	res2, err := conn.Receive()
	fmt.Printf("Receive res2:%v\n", res2)
	res3, err := conn.Receive()
	fmt.Printf("Receive res3:%s\n", res3)
}

//multi 事务
func transDemo() {
	conn, err := redis.Dial("tcp", "10.1.210.69:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()

	//MULTI, EXEC, DISCARD 和 WATCH 是构成 Redis 事务的基础。
	//MULTI：开启事务
	//EXEC：执行事务
	//DISCARD：取消事务
	//WATCH：监视事务中的键变化，一旦有改变则取消事务。
	conn.Send("MULTI")
	conn.Send("INCR", "foo")
	conn.Send("INCR", "bar")
	r, err := conn.Do("EXEC")
	fmt.Println(r)
}
