package main

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.TODO()
	// r := NewRedis("localhost:6379", "", "1234")
	r := NewRedis2()

	key := "chave"
	valueSet := "valor de teste"

	r.Set(ctx, key, valueSet, time.Second)

	valueReturn, _ := r.Get(ctx, key)

	fmt.Println(valueReturn)
}

type Redis struct {
	ClientCluster *redis.ClusterClient
	Client        *redis.Client
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

func NewRedis2() *Redis {
	return &Redis{
		ClientCluster: redis.NewClusterClient(
			&redis.ClusterOptions{
				Addrs: []string{"127.0.0.1:23000"},
			},
		),
	}
}

func (r *Redis) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	if r.ClientCluster != nil {
		return r.ClientCluster.Set(ctx, key, value, ttl).Err()
	} else {
		return r.Client.Set(ctx, key, value, ttl).Err()
	}
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	if r.ClientCluster != nil {
		return r.ClientCluster.Get(ctx, key).Result()
	} else {
		return r.Client.Get(ctx, key).Result()
	}
}
