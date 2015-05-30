package redis

import (
	"fmt"
	"github.com/gogather/com/log"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func Test_redis(t *testing.T) {
	var c Cache
	c.Open("tcp", "127.0.0.1", 6379)

	var data User
	data.Name = "李俊"
	data.Age = 25

	var to1 User
	log.Blueln("=== test hset and hget ===")
	err := c.Hset("key", "key2", data)
	if err == nil {
		err := c.Hget("key", "key2", &to1)
		if err == nil {
			log.Greenln(to1)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

	// del
	log.Blueln("=== test del ===")
	c.Del("key")
	var to2 User
	err = c.Hget("key", "key2", &to2)
	if err == nil {
		log.Greenln(to2)
	} else {
		fmt.Println(err)
	}

	// del again
	log.Blueln("=== test del again ===")
	c.Del("key")
	var to3 User
	err = c.Hget("key", "key2", &to3)
	if err == nil {
		log.Greenln(to3)
	} else {
		fmt.Println(err)
	}

	// test hdel
	var to4 User
	log.Blueln("=== test hdel ===")
	err = c.Hset("key", "key2", data)
	if err == nil {
		err := c.Hget("key", "key2", &to4)
		if err == nil {
			log.Greenln(to4)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

	// del
	c.Hdel("key", "key2")
	var to5 User
	err = c.Hget("key", "key2", &to5)
	if err == nil {
		log.Greenln(to5)
	} else {
		fmt.Println(err)
	}

}
