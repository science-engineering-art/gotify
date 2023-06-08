package persistence

import (
	"context"
	"encoding/base64"
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
	b64 := base64.RawStdEncoding.EncodeToString(key)

	err := rdb.Set(context.TODO(), b64, *data, time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rdb *RedisDb) Read(key []byte, start int32, end int32) (data *[]byte, err error) {
	b64 := base64.RawStdEncoding.EncodeToString(key)

	val, err := rdb.Get(context.TODO(), b64).Result()
	if err != nil {
		return nil, err
	}

	*data, err = base64.RawStdEncoding.DecodeString(val)
	return
}

func (rdb *RedisDb) Delete(key []byte) error {
	b64 := base64.RawStdEncoding.EncodeToString(key)

	_, err := rdb.Client.Del(context.TODO(), b64).Result()

	return err
}
