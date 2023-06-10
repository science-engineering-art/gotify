package persistence

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/jbenet/go-base58"
	"github.com/redis/go-redis/v9"
)

type RedisDb struct {
	redis.Client
}

func NewRedisDb(ip string) *RedisDb {
	addr := fmt.Sprintf("%s:6379", ip)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &RedisDb{*rdb}
}

func (rdb *RedisDb) Create(key []byte, data *[]byte) error {
	fmt.Printf("INIT RedisDb.Create(): %v with len(data)=%d\n", key, len(*data))
	defer fmt.Printf("EXIT RedisDb.Create(): %v\n", key)

	keyB64 := base58.Encode(key)
	dataB64 := base58.Encode(*data)

	err := rdb.Set(context.TODO(), keyB64, dataB64, time.Hour).Err()
	if err != nil {
		fmt.Println("ERROR rdb.Set()")
		return err
	}

	saved, _ := rdb.Read(key, 0, 0)

	if bytes.Equal(*data, *saved) {
		fmt.Printf("OKKK rdb.Set(%s) & len(saved)=%d\n", keyB64, len(*saved))
	} else {
		fmt.Println("ERROR rdb.Set()")
	}

	return nil
}

func (rdb *RedisDb) Read(key []byte, start int32, end int32) (*[]byte, error) {
	fmt.Println("INIT RedisDb.Read()")
	defer fmt.Println("EXIT RedisDb.Read()")

	keyB64 := base58.Encode(key)

	dataB64, err := rdb.Get(context.TODO(), keyB64).Result()
	if err != nil {
		fmt.Println("ERROR rdb.Get()")
		return nil, err
	}
	data := base58.Decode(dataB64)

	if end == 0 || end > int32(len(data)) {
		end = int32(len(data))
	}

	resp := data[start:end]

	fmt.Println("The resquested data is", resp)

	return &resp, err
}

func (rdb *RedisDb) Delete(key []byte) error {
	b64 := base58.Encode(key)

	_, err := rdb.Client.Del(context.TODO(), b64).Result()

	return err
}
