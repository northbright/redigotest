package main

import (
	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool

	replyTestCmds = []Command{
		Command{Cmd: "FLUSHALL", Args: []interface{}{}},
		Command{Cmd: "INFO", Args: []interface{}{"keyspace"}},
		Command{Cmd: "SET", Args: []interface{}{"student:num", 0}},
		Command{Cmd: "GET", Args: []interface{}{"student:num"}},
		Command{Cmd: "HSET", Args: []interface{}{"student:1", "name", "Bob"}},
		Command{Cmd: "HSET", Args: []interface{}{"student:1", "age", "20"}},
		Command{Cmd: "HGET", Args: []interface{}{"student:1", "name"}},
		Command{Cmd: "INCR", Args: []interface{}{"student:num"}},
		Command{Cmd: "GET", Args: []interface{}{"student:num"}},
		Command{Cmd: "LPUSH", Args: []interface{}{"laststudents", "1", "2", "3"}},
		Command{Cmd: "LRANGE", Args: []interface{}{"laststudents", 0, -1}},
		Command{Cmd: "HSCAN", Args: []interface{}{"student:1", 0}},
	}
)

func main() {
	pool = NewRedisPool(":6379", false, "")
	c := pool.Get()
	defer c.Close()
	// Test reply types of Golang for Redis Commands.
	GetReplyType(c, replyTestCmds)
}
