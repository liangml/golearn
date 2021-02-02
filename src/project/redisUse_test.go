package project

import (
	"fmt"
	"testing"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func TestRedisPool(t *testing.T) {
	// 从pool中取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "name", "tom猫")
	if err != nil {
		fmt.Println("set失败")
		return
	}
	// 取出
	r, err := redis.String(conn.Do("get", "name"))
	t.Log(r)

}

func TestRedisUse(t *testing.T) {
	// 链接redis数据库
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("链接失败")
		return
	}
	defer conn.Close() // 关闭链接
	// 写入string
	_, err = conn.Do("set", "name", "jack")
	if err != nil {
		fmt.Println("set失败")
		return
	}
	// 读取string
	// r 为interface，需要转换
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("读取失败")
		return
	}

	// hash类型
	_, err = conn.Do("hset", "user01", "name", "tom")
	_, err = conn.Do("hset", "user01", "age", "10")
	hrs, err := redis.String(conn.Do("hget", "user01", "name"))
	hri, err := redis.Int(conn.Do("hget", "user01", "age"))

	t.Log(r, hrs, hri)
	fmt.Println("操作成功")
}
