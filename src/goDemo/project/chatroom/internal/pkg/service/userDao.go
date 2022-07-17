package service

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"goDemo/project/chatroom/internal/pkg/model"
)

// UserDao 定义一个 UserDao 结构体体, 完成对 User 结构体的各种操作.
type UserDao struct {
	pool *redis.Pool
}

// NewUserDao 使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool: pool,
	}
	return userDao
}

//根据用户 id 返回 一个 User实例+err
func (userDao *UserDao) getUserById(conn redis.Conn, id int) (user *model.User, err error) {

	//通过给定id 去 redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil { //表示在 users 哈希中，没有找到对应id
			err = model.ErrorUserNotexists
		}
		return
	}

	//把 res 反序列化成 User 实例
	user = &model.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// Login 完成登录的校验
func (userDao *UserDao) Login(userId int, userPwd string) (user *model.User, err error) {

	//先从 UserDao 的连接池中取出一根连接
	conn := userDao.pool.Get()
	defer conn.Close()

	user, err = userDao.getUserById(conn, userId)
	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = model.ErrorUserPwd
		return
	}
	return
}

func (userDao *UserDao) Register(user *model.User) (err error) {

	//先从UserDao 的连接池中取出一根连接
	conn := userDao.pool.Get()
	defer conn.Close()

	_, err = userDao.getUserById(conn, user.UserId)
	if err == nil {
		err = model.ErrorUserExists
		return
	}

	data, err := json.Marshal(user) //序列化
	if err != nil {
		return
	}

	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}
	return
}
