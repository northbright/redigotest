package main

import (
	"fmt"
)

// GetReplyType() does test Redis commands to see what's the reply type in Golang for each Redis command.
func GetReplyType(cmds []Command) {
	conn := pool.Get()
	defer conn.Close()

	// Do commands.
	fmt.Printf("GetReplyType(): do cmds\n")
	for _, v := range cmds {
		ret, err := conn.Do(v.Cmd, v.Args...)
		if err != nil {
			fmt.Printf("%v error, args: %v, error: %v\n", v.Cmd, v.Args, err)
		} else {
			fmt.Printf("%v ok, args: %v, ret type: %T, ret: %v\n", v.Cmd, v.Args, ret, ret)
		}
	}
}
