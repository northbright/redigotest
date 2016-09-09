# redigotest
redigotest contains test examples for using Redigo.

#### Get redigo reply types in Golang
It'll use GetReplyType() with test Redis commands to see what's the reply type in Golang for each Redis command.

    GetReplyType(): do cmds
    FLUSHALL ok, args: [], ret type: string, ret: OK
    SET ok, args: [student:num 0], ret type: string, ret: OK
    GET ok, args: [student:num], ret type: []uint8, ret: [48]
    HSET ok, args: [student:1 name Bob], ret type: int64, ret: 1
    HGET ok, args: [student:1 name], ret type: []uint8, ret: [66 111 98]
    INCR ok, args: [student:num], ret type: int64, ret: 1
    GET ok, args: [student:num], ret type: []uint8, ret: [49]
