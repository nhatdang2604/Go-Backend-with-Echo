package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang/glog"
	"github.com/nhatdang2604/Go-Backend-with-Echo/user_manager/models"
)

type RedisCache struct {
	Host    string
	DB      int
	Expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) ICache {

	//glog.Infof("host = %v\r\n db = %v\r\n exp = %v\r\n", host, db, exp)

	return &RedisCache{
		Host:    host,
		DB:      db,
		Expires: exp,
	}
}

func (cache *RedisCache) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.Host,
		Password: "",
		DB:       cache.DB,
	})
}

func (cache *RedisCache) Set(id int32, value interface{}) {
	client := cache.GetClient()

	json, err := json.Marshal(value)
	if nil != err {
		glog.Errorf("Error on setting to cache value: %v", err)
		return
	}

	client.Set(string(id), json, cache.Expires)
}

func (cache *RedisCache) Get(id int32) interface{} {
	client := cache.GetClient()

	val, err := client.Get(string(id)).Result()
	if nil != err {
		glog.Errorf("Error on getting the cache value: %v", err)
		return nil
	}

	user := models.User{}
	err = json.Unmarshal([]byte(val), &user)

	if nil != err {
		glog.Errorf("Error on unmarshalling while getting the cache value: %v", err)
		return nil
	}

	return &user
}
