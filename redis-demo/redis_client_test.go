package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var rds *redis.Client

func init() {
	rds = redis.NewClient(&redis.Options{
		Addr: "centos1:6379",
	})
}

func TestGetAndSet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	key := "test-key"
	if err := rds.Get(ctx, key).Err(); err != redis.Nil {
		t.Fatalf("it shouldn't exist for key: %s, err: %v", key, err)
	}

	value := "test-value"
	rds.Set(ctx, key, value, 3*time.Second)
	if val := rds.Get(ctx, key).Val(); val != value {
		t.Fatalf("it expects to get value: %s, but it got value: %s", value, val)
	}

	rds.Del(ctx, key)
}

func TestHashGetAndSet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	key := "hash-key"
	if err := rds.Exists(ctx, key).Err(); err != redis.Nil {
		t.Fatalf("it shouldn't exist for key: %s, err: %v", key, err)
	}

	itemKey := "hash-itemKey"
	itemVal := "hash-itemVal"

	rds.HSet(ctx, key, itemKey, itemVal)
	if val := rds.HGet(ctx, key, itemKey).Val(); val != itemVal {
		t.Fatalf("it expects to get value: %s, but it got value: %s", itemVal, val)
	}

	rds.Del(ctx, key)
}
