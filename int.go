package redigotest

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

func SetInt(c redis.Conn, k string, v interface{}) {
	c.Do("DEL", k)
	log.Printf("SET k: %v, v: %v(%T)\n", k, v, v)
	c.Do("SET", k, v)
	ret, _ := redis.String(c.Do("DEBUG", "OBJECT", k))
	// Check encoding and value size.
	log.Printf("DEBUG OBJECT %v: %v\n", k, ret)
}
