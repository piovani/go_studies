package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.TODO()
	r := NewRedis("localhost:6379", "", "1234")

	key := "chave"
	valueSet := "valor de teste"

	r.Set(ctx, key, valueSet, time.Second)

	valueReturn, _ := r.Get(ctx, key)

	fmt.Println(valueReturn)
}

type Redis struct {
	Client *redis.Client
}

func NewRedis(host string, user string, password string) *Redis {
	return &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:     host,
			Username: user,
			Password: password,
		}),
	}
}

func (r *Redis) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}
