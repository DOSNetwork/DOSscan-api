package cache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/gomodule/redigo/redis"
)

func getConn() redis.Conn {
	pool := &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 3,
		// max number of connections
		MaxActive: 10,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	conn := pool.Get()
	// call Redis PING command to test connectivity
	s, err := redis.String(conn.Do("PING"))
	if err != nil {
		fmt.Printf("PING err = %s\n", err)
	}
	fmt.Printf("PING Response = %s\n", s)

	return conn
}

func TestCache(t *testing.T) {
	conn := getConn()
	c := NewCacheRepo(conn)

	if err := c.Set(context.Background(), "address+0x123", "123"); err != nil {
		t.Errorf("TestCache Error : Err %s ", err)
	}
	if value, err := c.Get(context.Background(), "address+0x123"); err != nil {
		t.Errorf("TestCache Error : Err %s ", err)
	} else {
		fmt.Println("result ", value)
	}
}

func TestExpire(t *testing.T) {
	conn := getConn()
	c := NewCacheRepo(conn)
	var logs []models.LogGrouping
	if err := c.Set(context.Background(), "address+0x123", "123"); err != nil {
		t.Errorf("TestCache Error : Err %s ", err)
	}
	time.Sleep(16 * time.Second)
	if value, err := c.Get(context.Background(), "address+0x123"); err != redis.ErrNil {
		t.Errorf("TestCache Error : Err %s ", err)
	} else {
		fmt.Println("result ", value)
	}
}
