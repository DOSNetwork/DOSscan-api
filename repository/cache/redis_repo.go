package cache

import (
	"context"

	"github.com/DOSNetwork/DOSscan-api/repository"

	"github.com/gomodule/redigo/redis"
)

type cacheRepo struct {
	conn redis.Conn
}

func NewCacheRepo(conn redis.Conn) repository.Cache {
	return &cacheRepo{
		conn: conn,
	}
}

func (c *cacheRepo) Set(ctx context.Context, key string, value string) (err error) {

	// SET object
	_, err = c.conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = c.conn.Do("EXPIRE", key, 15)

	return
}

func (c *cacheRepo) Get(ctx context.Context, key string) (value string, err error) {
	value, err = redis.String(c.conn.Do("GET", key))
	return
}
