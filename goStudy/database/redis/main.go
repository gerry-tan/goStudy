package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
)

func single() {
	conn, err := redis.Dial("tcp", "10.42.0.121:7003")
	if err != nil {
		fmt.Printf("connect failed, error: %v", err)
		return
	}

	defer conn.Close()

	_, err = conn.Do("Set", "abc", 100)
	conn.Do("expire", "abc", 10)
	if err != nil {
		fmt.Printf("set value failed, error: %v", err)
		return
	}

	time.Sleep(11 * time.Second)

	i, err := redis.Int(conn.Do("Get", "abc"))
	if err != nil {
		fmt.Printf("get key value failed, error: %v", err)
		return
	}

	fmt.Println(i)

	_, err = conn.Do("Del", "abc")
	if err != nil {
		fmt.Printf("del key failed, error: %v", err)
		return
	}
}

var pool *redis.Pool

func redisPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			//如果设置有秘钥
			/*if _, err := conn.Do("AUTH", password); err != nil {
				conn.Close()
				return nil, err
			}*/
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func doPool(wg *sync.WaitGroup) {
	pool = redisPool("10.42.0.121:7003", "")

	i := 0
	for i < 10 {
		time.Sleep(time.Second)
		conn := pool.Get()
		_, err := conn.Do("set", "abc", 123)
		if err != nil {
			fmt.Printf("set key failed, error: %v\n", err)
			continue
		}

		res, err := redis.Int(conn.Do("get", "abc"))
		if err != nil {
			fmt.Printf("get value failed, error: %v\n", err)
			continue
		}
		i++
		fmt.Printf("%d:\t%d\n", i, res)
	}
	wg.Done()
}

func main() {
	//go single()

	var wg sync.WaitGroup
	wg.Add(1)
	go doPool(&wg)
	wg.Wait()
}
