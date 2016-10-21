package redigotest_test

import (
	"github.com/northbright/redigotest"
)

// Example_GetReplyType tests the return value and types of Redis Commands.
func Example_GetReplyType() {
	cmds := []redigotest.Command{
		redigotest.Command{Cmd: "FLUSHALL", Args: []interface{}{}},
		redigotest.Command{Cmd: "INFO", Args: []interface{}{"keyspace"}},
		redigotest.Command{Cmd: "SET", Args: []interface{}{"student:num", 0}},
		redigotest.Command{Cmd: "GET", Args: []interface{}{"student:num"}},
		redigotest.Command{Cmd: "HSET", Args: []interface{}{"student:1", "name", "Bob"}},
		redigotest.Command{Cmd: "HSET", Args: []interface{}{"student:1", "age", "20"}},
		redigotest.Command{Cmd: "HGET", Args: []interface{}{"student:1", "name"}},
		redigotest.Command{Cmd: "INCR", Args: []interface{}{"student:num"}},
		redigotest.Command{Cmd: "GET", Args: []interface{}{"student:num"}},
		redigotest.Command{Cmd: "LPUSH", Args: []interface{}{"laststudents", "1", "2", "3"}},
		redigotest.Command{Cmd: "LRANGE", Args: []interface{}{"laststudents", 0, -1}},
		redigotest.Command{Cmd: "HSCAN", Args: []interface{}{"student:1", 0}},
	}

	pool := redigotest.NewRedisPool(":6379", false, "")
	c := pool.Get()
	defer c.Close()

	redigotest.GetReplyType(c, cmds)
	// Output:
}
