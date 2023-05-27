package redis

import (
	"testing"

	"github.com/giautm/httpcache/test"
	"github.com/gomodule/redigo/redis"
)

const testServer = "localhost:6379"

func TestRedisCache(t *testing.T) {
	conn, err := redis.Dial("tcp", testServer)
	if err != nil {
		// TODO: rather than skip the test, fall back to a faked redis server
		t.Skipf("skipping test; no server running at %s, err = %v", testServer, err)
	}
	conn.Do("FLUSHALL")

	test.Cache(t, NewWithClient(conn))
}
