package redigotest

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

// GetReplyType does test Redis commands to see what's the reply type in Golang for each Redis command.
func GetReplyType(c redis.Conn, cmds []Command) {
	// Do commands.
	log.Printf("GetReplyType(): do cmds\n")
	for _, v := range cmds {
		ret, err := c.Do(v.Cmd, v.Args...)
		if err != nil {
			log.Printf("%v error, args: %v, error: %v\n", v.Cmd, v.Args, err)
		} else {
			log.Printf("%v ok, args: %v, ret type: %T, ret: %v\n", v.Cmd, v.Args, ret, ret)
		}
	}
}
