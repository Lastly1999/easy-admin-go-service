package captcha

import (
	"context"
	"easy-admin-go-service/pkg/redis_store"
	"fmt"
	"time"
)

var ctx = context.Background()

const CAPTCHA = "base:admin:auth:captcha:"

type RedisStore struct {
}

// Set 实现获取captcha的方法
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	//time.Minute*2：有效时间2分钟
	err := redis_store.RedisDb.Set(ctx, key, value, time.Minute*2).Err()
	return err
}

// Get 实现获取captcha的方法
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := redis_store.RedisDb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		//clear为true，验证通过，删除这个验证码
		err := redis_store.RedisDb.Del(ctx, key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

// Verify 实现验证captcha的方法
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
