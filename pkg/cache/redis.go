package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-server/config"
	"strings"
	"time"
)

const (
	DefaultRedisConnection = "default"
)

var (
	rds = new(Redis)

	ctx = context.Background()
)

type Redis struct {
	// 链接集合
	Map map[string]*redis.Client
	// 当前链接
	Conn *redis.Client
}

func InitRedis() {
	conf := config.Get().Redis
	rds.Map = make(map[string]*redis.Client, len(conf))

	for key, value := range conf {
		conn := redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%d", value.Addr, value.Port),
			Password:     value.Password,
			DB:           value.Db,
			MaxRetries:   value.MaxRetries,   // 最大的重试次数
			PoolSize:     value.PoolSize,     // 连接池最大连接数，默认为 CPU 数 * 10
			PoolTimeout:  value.PoolTimeout,  // 连接池超时时间
			MinIdleConns: value.MinIdleConns, // 最小空闲连接数
			IdleTimeout:  value.IdleTimeout,  // 空闲连接超时时间
			DialTimeout:  value.DialTimeout,  // 建立连接超时时间
			ReadTimeout:  value.ReadTimeout,  // 读超时时间
			WriteTimeout: value.WriteTimeout, // 写超时时间
		})

		_, err := conn.Ping(ctx).Result()
		if err == nil {
			rds.Map[strings.ToLower(key)] = conn
		}
	}
}

// 获取链接
func GetRedis(connection string) *Redis {
	rds.Conn = rds.Map[strings.ToLower(connection)]
	return rds
}

// 获取默认链接
func GetDefaultRedis() *Redis {
	return GetRedis(DefaultRedisConnection)
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return r.Conn.Set(ctx, key, value, expiration).Result()
}

func (r *Redis) SetEX(key string, value interface{}, expiration time.Duration) (string, error) {
	return r.Conn.SetEX(ctx, key, value, expiration).Result()
}

func (r *Redis) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.Conn.SetNX(ctx, key, value, expiration).Result()
}

func (r *Redis) Exists(keys ...string) bool {
	v, err := r.Conn.Exists(ctx, keys...).Result()
	if (err != nil) || (v == 0) {
		return false
	}
	return true
}

func (r *Redis) GetSet(key string, value interface{}) *redis.StringCmd {
	return r.Conn.GetSet(ctx, key, value)
}

func (r *Redis) Get(key string) (string, error) {
	return r.Conn.Get(ctx, key).Result()
}

func (r *Redis) Expire(key string, expiration time.Duration) (bool, error) {
	return r.Conn.Expire(ctx, key, expiration).Result()
}

func (r *Redis) Del(keys ...string) (int64, error) {
	return r.Conn.Del(ctx, keys...).Result()
}

func (r *Redis) HSet(key string, values ...interface{}) (int64, error) {
	return r.Conn.HSet(ctx, key, values).Result()
}

func (r *Redis) HSetNX(key, field string, value interface{}) (bool, error) {
	return r.Conn.HSetNX(ctx, key, field, value).Result()
}

func (r *Redis) HMSet(key string, values ...interface{}) (bool, error) {
	return r.Conn.HMSet(ctx, key, values).Result()
}

func (r *Redis) HExists(key, field string) bool {
	v, err := r.Conn.HExists(ctx, key, field).Result()
	if (err != nil) || !v {
		return false
	}
	return true
}

func (r *Redis) HGet(key, field string) (string, error) {
	return r.Conn.HGet(ctx, key, field).Result()
}

func (r *Redis) HGetAll(key string) (map[string]string, error) {
	return r.Conn.HGetAll(ctx, key).Result()
}

func (r *Redis) HMGet(key string, fields ...string) (interface{}, error) {
	return r.Conn.HMGet(ctx, key, fields...).Result()
}

func (r *Redis) HKeys(key string) ([]string, error) {
	return r.Conn.HKeys(ctx, key).Result()
}

func (r *Redis) HVals(key string) ([]string, error) {
	return r.Conn.HVals(ctx, key).Result()
}

func (r *Redis) HLen(key string) (int64, error) {
	return r.Conn.HLen(ctx, key).Result()
}

func (r *Redis) HIncrBy(key, field string, incr int64) (int64, error) {
	return r.Conn.HIncrBy(ctx, key, field, incr).Result()
}

func (r *Redis) HIncrByFloat(key, field string, incr float64) (float64, error) {
	return r.Conn.HIncrByFloat(ctx, key, field, incr).Result()
}

func (r *Redis) HDel(key string, fields ...string) (int64, error) {
	return r.Conn.HDel(ctx, key, fields...).Result()
}

// 关闭链接
func Close(connection string) error {
	return rds.Map[strings.ToLower(connection)].Close()
}

// 关闭所有链接
func CloseAll() error {
	for _, client := range rds.Map {
		if err := client.Close(); err != nil {
			return err
		}
	}

	return nil
}
