package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const Timeout = 1

type Cache struct {
	Client *redis.Client
}

func NewCache() *Cache {
	rdb := redis.NewClient(&redis.Options{})
	return &Cache{Client: rdb}
}

func (c *Cache) Start(addr, password string) (*Cache, error) {
	c.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	err := c.isConnected()
	if err != nil {
		log.Fatalf("cache ping err: %s\n", err)
		return nil, err
	}
	return c, nil
}

func (c *Cache) isConnected() error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()
	_, err := c.Client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Get(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	strValue, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(strValue), value)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Set(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()
	bData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	status := c.Client.Set(ctx, key, bData, 0)
	if status.Err() != nil {
		return err
	}
	return nil
}

func (c *Cache) Subscribe(channels []string) error {
	err := c.isConnected()
	if err != nil {
		log.Printf("Encountered an error while subscribing: %s\n", err)
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()
	c.Client.Subscribe(ctx, channels...)
	return nil
}

func (c *Cache) Receive(channel string) (<-chan *redis.Message, error) {
	err := c.isConnected()
	if err != nil {
		log.Printf("Encountered an error while subscribing: %s\n", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()
	return c.Client.Subscribe(ctx, channel).Channel(), nil
}

func (c *Cache) Publish(channel string, message string) error {
	err := c.isConnected()
	if err != nil {
		log.Printf("Encountered an error while subscribing: %s\n", err)
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	c.Client.Publish(ctx, channel, message)
	return nil
}
