package main

import (
	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool

	replyTestCmds = []Command{
		Command{Cmd: "FLUSHALL", Args: []interface{}{}},
		Command{Cmd: "SET", Args: []interface{}{"student:num", 0}},
		Command{Cmd: "GET", Args: []interface{}{"student:num"}},
		Command{Cmd: "HSET", Args: []interface{}{"student:1", "name", "Bob"}},
		Command{Cmd: "HGET", Args: []interface{}{"student:1", "name"}},
		Command{Cmd: "INCR", Args: []interface{}{"student:num"}},
		Command{Cmd: "GET", Args: []interface{}{"student:num"}},
	}
)

func main() {
	pool = NewRedisPool(":6379", "123456")
	// Test reply types of Golang for Redis Commands.
	GetReplyType(replyTestCmds)
}
