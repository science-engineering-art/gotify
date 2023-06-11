package persistence

import (
	"bytes"
	"context"
	"errors"
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
	//fmt.Printf("INIT RedisDb.Create(%v): len(data)=%d\n", key, len(*data))
	// defer //fmt.Printf("EXIT RedisDb.Create(%s)\n", key)

	//fmt.Println("Before Key Encode")
	keyB58 := base58.Encode(key)
	//fmt.Println("After Key Encode")

	// //fmt.Println("Before *data Encode")
	// dataB58 := base58.Encode(*data)
	// //fmt.Println("After *data Encode")

	//fmt.Println("KEY ID", keyB58)

	//fmt.Printf("Before rdb.Set(%s)\n", keyB58)
	err := rdb.Set(context.TODO(), keyB58, *data, time.Hour).Err()
	//fmt.Printf("After rdb.Set(%s)\n", keyB58)
	if err != nil {
		//fmt.Println("ERROR line:36 rdb.Set() err:", err)
		return err
	}

	//fmt.Println("Before Check")
	saved, _ := rdb.Read(key, 0, 0)
	//fmt.Println("After Check")
	if saved == nil {
		//fmt.Println("ERROR line:43 not created")
		return errors.New("not created")
	}

	if bytes.Equal(*data, *saved) {
		//fmt.Printf("OKKK rdb.Set(%s) & len(saved)=%d\n", keyB58, len(*saved))
	} else {
		//fmt.Println("ERROR line:45 rdb.Set()")
	}

	return nil
}

func (rdb *RedisDb) Read(key []byte, start int64, end int64) (*[]byte, error) {
	keyB58 := base58.Encode(key)

	//fmt.Printf("INIT RedisDb.Read(%s)\n", keyB58)
	// defer //fmt.Printf("EXIT RedisDb.Read(%s)\n", keyB58)

	dataB58, err := rdb.Get(context.TODO(), keyB58).Result()
	if err != nil {
		//fmt.Printf("ERROR line:64 rdb.Get(%s)\n", keyB58)
		return nil, err
	}
	data := []byte(dataB58)

	if end == 0 || end > int64(len(data)) {
		end = int64(len(data))
	}

	resp := data[start:end]

	//fmt.Println("The resquested data is", resp)

	return &resp, err
}

func (rdb *RedisDb) Delete(key []byte) error {
	b58 := base58.Encode(key)

	_, err := rdb.Client.Del(context.TODO(), b58).Result()

	return err
}

func (rdb *RedisDb) GetKeys() [][]byte {
	keys, err := rdb.Keys(context.TODO(), "*").Result()
	result := [][]byte{}
	if err != nil {
		return result
	}
	for _, key := range keys {
		keyBytes := base58.Decode(key)
		result = append(result, keyBytes)
	}
	return result
}
