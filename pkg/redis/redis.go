package redis

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

// Client represents Redis client
type Client struct {
	c *redis.Client
}

func new(addr, pw string, port int) (*Client, error) {
	opts := redis.Options{
		Addr: addr + ":" + strconv.Itoa(port),
	}

	if pw != "" {
		opts.Password = pw
	}

	client := redis.NewClient(&opts)

	_, err := client.Ping().Result()

	if err != nil {
		return nil, fmt.Errorf("Cannot connect to redis: %v", err)
	}

	return &Client{c: client}, nil
}
