# redigotest

[![Build Status](https://travis-ci.org/northbright/redigotest.svg?branch=master)](https://travis-ci.org/northbright/redigotest)

redigotest contains test examples for using [redigo](https://github.com/garyburd/redigo).

#### Get type of reply after [exectuing commands](https://godoc.org/github.com/garyburd/redigo/redis#hdr-Executing_Commands)
It'll use GetReplyType() with test Redis commands to see what's the reply type in Golang for each Redis command.

    GetReplyType(): do cmds
    FLUSHALL ok, args: [], ret type: string, ret: OK
    SET ok, args: [student:num 0], ret type: string, ret: OK
    GET ok, args: [student:num], ret type: []uint8, ret: [48]
    HSET ok, args: [student:1 name Bob], ret type: int64, ret: 1
    HGET ok, args: [student:1 name], ret type: []uint8, ret: [66 111 98]
    INCR ok, args: [student:num], ret type: int64, ret: 1
    GET ok, args: [student:num], ret type: []uint8, ret: [49]
    LPUSH ok, args: [laststudents 1 2 3], ret type: int64, ret: 3
    LRANGE ok, args: [laststudents 0 -1], ret type: []interface {}, ret: [[51] [50] [49]]

