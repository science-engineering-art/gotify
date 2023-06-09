package persistence

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisDb struct {
	redis.Client
}

func NewRedisDb() *RedisDb {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &RedisDb{*rdb}
}

func (rdb *RedisDb) Create(key []byte, data *[]byte) error {
	keyB64 := base64.RawStdEncoding.EncodeToString(key)
	dataB64 := base64.RawStdEncoding.EncodeToString(*data)

	err := rdb.Set(context.TODO(), keyB64, dataB64, time.Minute).Err()
	if err != nil {
		return err
	}
	fmt.Println("Store Data with key:", keyB64)

	return nil
}

func (rdb *RedisDb) Read(key []byte, start int32, end int32) (*[]byte, error) {
	keyB64 := base64.RawStdEncoding.EncodeToString(key)

	dataB64, err := rdb.Get(context.TODO(), keyB64).Result()
	if err != nil {
		return nil, err
	}
	data, err := base64.RawStdEncoding.DecodeString(dataB64)

	if end == 0 || end > int32(len(data)) {
		end = int32(len(data))
	}

	resp := data[start:end]

	fmt.Println("The resquested data is", resp)

	return &resp, err
}

func (rdb *RedisDb) Delete(key []byte) error {
	b64 := base64.RawStdEncoding.EncodeToString(key)

	_, err := rdb.Client.Del(context.TODO(), b64).Result()

	return err
}
