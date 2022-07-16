package redis

import (
	"fmt"
	"gopkg.in/redis.v4"
	"sync"
	"time"
)

/*
使用 https://github.com/go-redis/r... 客户端, 安装方式如下:
go get gopkg.in/redis.v4
*/

// 创建 redis 客户端
func createClient() *redis.Client {
	//该方法接收一个 redis.Options 对象参数，通过该参数可以配置 redis 相关属性，如 redis 服务器地址, 数据库名, 数据库密码等
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

// String 操作
func stringOperation(client *redis.Client) {
	// 第三个参数是过期时间, 如果是 0, 则表示没有过期时间.
	err := client.Set("name", "xys", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)

	err = client.Set("age", "20", 1*time.Second).Err()
	if err != nil {
		panic(err)
	}

	client.Incr("age") // 自增
	client.Incr("age") // 自增
	client.Decr("age") // 自减

	val, err = client.Get("age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("age", val) // age 的值为21

	// 因为 key "age" 的过期时间是一秒钟, 因此当一秒后, 此 key 会自动被删除了.
	time.Sleep(1 * time.Second)
	val, err = client.Get("age").Result()
	if err != nil {
		// 因为 key "age" 已经过期了, 因此会有一个 redis: nil 的错误.
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("age", val)
}

// list 操作
func listOperation(client *redis.Client) {

	//在名称为 fruit 的 list 尾和头添加一个值为 value 的元素
	client.RPush("fruit", "apple")
	client.LPush("fruit", "banana")
	length, err := client.LLen("fruit").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("length: ", length) // 长度为2

	//返回并删除名称为 fruit 的list中的首元素
	value, err := client.LPop("fruit").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fruit: ", value)

	value, err = client.RPop("fruit").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fruit: ", value)
}

// set 操作
func setOperation(client *redis.Client) {
	/*
		sadd(key, member)：向名称为key的set中添加元素member
		srem(key, member) ：删除名称为key的set中的元素member
		spop(key) ：随机返回并删除名称为key的set中一个元素
		smove(srckey, dstkey, member) ：移到集合元素
		scard(key) ：返回名称为key的set的基数
		sismember(key, member) ：member是否是名称为key的set的元素
		sinter(key1, key2,…key N) ：求交集
		sinterstore(dstkey, (keys)) ：求交集并将交集保存到dstkey的集合
		sunion(key1, (keys)) ：求并集
		sunionstore(dstkey, (keys)) ：求并集并将并集保存到dstkey的集合
		sdiff(key1, (keys)) ：求差集
		sdiffstore(dstkey, (keys)) ：求差集并将差集保存到dstkey的集合
		smembers(key) ：返回名称为key的set的所有元素
		srandmember(key) ：随机返回名称为key的set的一个元素
	*/
	client.SAdd("blacklist", "Obama")
	client.SAdd("blacklist", "Elder")
	client.SAdd("whitelist", "Elder")

	// 判断元素是否在集合中
	isMember, err := client.SIsMember("blacklist", "Bush").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Is Bush in blacklist: ", isMember)

	// 求交集, 即既在黑名单中, 又在白名单中的元素
	names, err := client.SInter("blacklist", "whitelist").Result()
	if err != nil {
		panic(err)
	}
	// 获取到的元素是 "Elder"
	fmt.Println("Inter result: ", names)

	// 获取指定集合的所有元素
	all, err := client.SMembers("blacklist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("All member: ", all)
}

// hash 操作
func hashOperation(client *redis.Client) {
	client.HSet("user_xys", "name", "xys")
	client.HSet("user_xys", "age", "18")

	// 批量地向 user_test 的 hash 中添加元素 name 和 age
	client.HMSet("user_test", map[string]string{"name": "test", "age": "20"})

	// 批量获取 user_test 的 hash 中的指定字段的值.
	fields, err := client.HMGet("user_test", "name", "age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fields in user_test: ", fields)

	// 获取 user_xys 的 hash 中的字段个数
	length, err := client.HLen("user_xys").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("field count in user_xys: ", length) // 字段个数为2

	// 删除 user_test 的 age 字段
	client.HDel("user_test", "age")
	age, err := client.HGet("user_test", "age").Result()
	if err != nil {
		fmt.Printf("Get user_test age error: %v\n", err)
	} else {
		fmt.Println("user_test age is: ", age) // 字段个数为2
	}
}

// 创建 redis 连接池客户端
func poolClient() *redis.Client {
	//redis.v4 包实现了 redis 的连接池管理, 因此就不需要自己手动管理 redis 连接了。
	//默认情况下, redis.v4 的 redis 连接池大小是10, 不过可以在初始化 redis 客户端时自行设置连接池的大小
	//通过 redis.Options 的 PoolSize 属性, 我们设置了 redis 连接池的大小为 5
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 5,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

// redis.v4 的连接池管理
func connectPool(client *redis.Client) {
	/*
		该例子启动了10个 routine 来不断向 redis 读写数据, 然后通过 client.PoolStats() 获取连接池的信息. 运行输出如下:
		PoolStats, TotalConns: 5, FreeConns: 1
		PoolStats, TotalConns: 5, FreeConns: 1
		PoolStats, TotalConns: 5, FreeConns: 1
		PoolStats, TotalConns: 5, FreeConns: 1
		PoolStats, TotalConns: 5, FreeConns: 1
		PoolStats, TotalConns: 5, FreeConns: 2
		PoolStats, TotalConns: 5, FreeConns: 2
		PoolStats, TotalConns: 5, FreeConns: 3
		PoolStats, TotalConns: 5, FreeConns: 4
		PoolStats, TotalConns: 5, FreeConns: 5

		通过输出可以看到, 此时最大的连接池数量确实是 5, 并且一开始时, 因为 goroutine 的数量大于 5, 会造成 redis 连接不足的情况。
		反映在 FreeConns 上就是前几次的输出 FreeConns 一直是 1), 当某个 goroutine 结束后, 会释放此 redis 连接, 因此 FreeConns 会增加
	*/
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
				client.Get(fmt.Sprintf("name%d", j)).Result()
			}
			fmt.Printf("PoolStats, TotalConns: %d, FreeConns: %d\n", client.PoolStats().TotalConns, client.PoolStats().FreeConns)
		}()
	}

	wg.Wait()
}

func testMain() {
	client := createClient()
	defer client.Close()

	stringOperation(client)
	listOperation(client)
	setOperation(client)
	hashOperation(client)

	client = poolClient()
	defer client.Close()
	connectPool(client)
}
