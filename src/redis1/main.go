package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	option := redis.DialPassword("123456")
	conn, err := redis.Dial("tcp", "192.168.0.26:6379", option)
	if err != nil {
		fmt.Println("redis connection failed..")
		return
	}

	defer conn.Close()

	_, err = conn.Do("Set", "name", "songzeceng") // 参数列表：指令、键、值
	if err != nil {
		fmt.Println("redis set failed..")
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	fmt.Println("result = ", r)

	_, err = conn.Do("HSet", "userhash01", "name", "szc") // 操作、哈希名、键、值
	_, err = conn.Do("HSet", "userhash01", "age", 23)

	r, err = redis.String(conn.Do("HGet", "userhash01", "name")) // 从哈希userhash01中读取name
	fmt.Println("hash name = ", r)
	age, err := redis.Int(conn.Do("HGet", "userhash01", "age")) // 读取int
	fmt.Println("hash age = ", age)

	_, err = conn.Do("MSet", "name", "songzeceng", "home", "Henan,Anyang")
	multi_r, err := redis.Strings(conn.Do("MGet", "name", "home")) // 注意String多了个s，而且multi_r是[]string
	fmt.Println(multi_r)
}
