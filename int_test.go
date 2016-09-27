package redigotest_test

import (
	"github.com/northbright/redigotest"
)

func Example_SetInt() {
	var vUint64 uint64 = 12345678
	var vInt64 int64 = 12345678
	var vStr string = "12345678"

	vArr := []interface{}{vUint64, vInt64, vStr}

	pool := redigotest.NewRedisPool(":6379", false, "")
	c := pool.Get()
	defer c.Close()

	for _, v := range vArr {
		redigotest.SetInt(c, "int-test", v)
	}
	// Output:
}
